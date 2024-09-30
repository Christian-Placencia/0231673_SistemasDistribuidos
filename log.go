package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	api "github.com/Christian-Placencia/0231673_SistemasDistribuidos/api/v1"
)

type Log struct {
	Dir    string
	Config Config

	activeSegment *Segment
	segments      []*Segment
}

// NewLog creates a new log, initializing the log by loading segments from disk or creating new ones.
func NewLog(dir string, c Config) (*Log, error) {
	l := &Log{
		Dir:    dir,
		Config: c,
	}

	return l, l.setup()
}

// setup initializes the log by loading existing segments from disk or creating a new segment.
func (l *Log) setup() error {
	files, err := os.ReadDir(l.Dir)
	if err != nil {
		return err
	}

	var baseOffsets []uint64
	for _, file := range files {
		off, isValid, err := parseBaseOffset(file.Name())
		if err != nil || !isValid {
			continue
		}
		baseOffsets = append(baseOffsets, off)
	}

	sort.Slice(baseOffsets, func(i, j int) bool { return baseOffsets[i] < baseOffsets[j] })

	// Load existing segments
	for _, off := range baseOffsets {
		if err := l.newSegment(off); err != nil {
			return err
		}
	}

	// Create new segment if no segments exist
	if l.segments == nil {
		if err := l.newSegment(l.Config.Segment.InitialOffset); err != nil {
			return err
		}
	}

	return nil
}

// newSegment creates a new segment in the log.
func (l *Log) newSegment(baseOffset uint64) error {
	s, err := newSegment(l.Dir, baseOffset, l.Config)
	if err != nil {
		return err
	}
	l.segments = append(l.segments, s)
	l.activeSegment = s
	return nil
}

// Append adds a new record to the active segment and returns the offset of the record.
func (l *Log) Append(record *api.Record) (uint64, error) {
	if l.activeSegment.IsMaxed() {
		if err := l.newSegment(l.activeSegment.nextOffset); err != nil {
			return 0, err
		}
	}
	return l.activeSegment.Append(record)
}

// Read retrieves a record by its offset from the appropriate segment.
func (l *Log) Read(off uint64) (*api.Record, error) {
	for _, s := range l.segments {
		if off >= s.baseOffset && off < s.nextOffset {
			return s.Read(off)
		}
	}
	return nil, io.EOF
}

// Close closes all segments in the log.
func (l *Log) Close() error {
	for _, s := range l.segments {
		if err := s.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Remove deletes the log directory and its contents.
func (l *Log) Remove() error {
	if err := l.Close(); err != nil {
		return err
	}
	return os.RemoveAll(l.Dir)
}

// Reset removes all existing segments and reinitializes the log.
func (l *Log) Reset() error {
	if err := l.Remove(); err != nil {
		return err
	}
	return l.setup()
}

// LowestOffset returns the offset of the oldest record in the log.
func (l *Log) LowestOffset() (uint64, error) {
	if len(l.segments) == 0 {
		return 0, fmt.Errorf("no segments found")
	}
	return l.segments[0].baseOffset, nil
}

// HighestOffset returns the offset of the most recent record in the log.
func (l *Log) HighestOffset() (uint64, error) {
	if len(l.segments) == 0 {
		return 0, fmt.Errorf("no segments found")
	}
	lastSegment := l.segments[len(l.segments)-1]
	return lastSegment.nextOffset - 1, nil
}

// Reader returns an io.Reader that reads all records across the segments.
// Reader returns an io.Reader that reads all records across the segments.
func (l *Log) Reader() io.Reader {
	readers := make([]io.Reader, len(l.segments))
	for i, segment := range l.segments {
		readers[i] = &storeReader{store: segment.store}
	}
	return io.MultiReader(readers...)
}

// Truncate removes all segments whose highest offset is below the provided offset.
func (l *Log) Truncate(lowestOffset uint64) error {
	var remainingSegments []*Segment
	for _, s := range l.segments {
		if s.nextOffset <= lowestOffset {
			if err := s.Remove(); err != nil {
				return err
			}
		} else {
			remainingSegments = append(remainingSegments, s)
		}
	}
	l.segments = remainingSegments
	return nil
}

func parseBaseOffset(filename string) (offset uint64, isValid bool, err error) {
	// Extract the base name of the file without its extension
	base := filepath.Base(filename)

	// Split the filename at the period to separate the base offset from the extension
	parts := strings.Split(base, ".")
	if len(parts) != 2 {
		return 0, false, fmt.Errorf("invalid file name format: %s", filename)
	}

	// Ensure the file has a valid extension: ".store" or ".index"
	ext := parts[1]
	if ext != "store" && ext != "index" {
		return 0, false, fmt.Errorf("invalid file extension: %s", ext)
	}

	// Attempt to parse the base offset (the numeric part before the extension)
	offset, err = strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return 0, false, fmt.Errorf("invalid base offset: %v", err)
	}

	// If all parsing is successful, return the parsed offset and mark as valid
	return offset, true, nil
}
