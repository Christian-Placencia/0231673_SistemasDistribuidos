// log.go
package log

import (
	"io"
	"os"

	api "github.com/Robinthatdoesnotsuck/api/log_v1"
)

type Log struct {
	Dir    string
	Config Config

	activeSegment *Segment
	segments      []*Segment
}

func NewLog(dir string, c Config) (*Log, error) {
	l := &Log{
		Dir:    dir,
		Config: c,
	}

	return l, l.setup()
}

func (l *Log) setup() error {
	files, err := os.ReadDir(l.Dir)
	if err != nil {
		return err
	}

	var baseOffsets []uint64
	for _, file := range files {
		off, _, _ := parseBaseOffset(file.Name())
		baseOffsets = append(baseOffsets, off)
	}

	// Sort the base offsets in ascending order
	for _, off := range baseOffsets {
		if err := l.newSegment(off); err != nil {
			return err
		}
	}

	if l.segments == nil {
		if err := l.newSegment(l.Config.Segment.InitialOffset); err != nil {
			return err
		}
	}

	return nil
}

func (l *Log) newSegment(baseOffset uint64) error {
	s, err := newSegment(l.Dir, baseOffset, l.Config)
	if err != nil {
		return err
	}
	l.segments = append(l.segments, s)
	l.activeSegment = s
	return nil
}

func (l *Log) Append(record *api.Record) (uint64, error) {
	if l.activeSegment.IsMaxed() {
		if err := l.newSegment(l.activeSegment.nextOffset); err != nil {
			return 0, err
		}
	}

	return l.activeSegment.Append(record)
}

func (l *Log) Read(off uint64) (*api.Record, error) {
	for _, s := range l.segments {
		if off >= s.baseOffset && off < s.nextOffset {
			return s.Read(off)
		}
	}
	return nil, io.EOF
}

func (l *Log) Close() error {
	for _, s := range l.segments {
		if err := s.Close(); err != nil {
			return err
		}
	}
	return nil
}

func (l *Log) Remove() error {
	if err := l.Close(); err != nil {
		return err
	}
	return os.RemoveAll(l.Dir)
}

func (l *Log) Reset() error {
	if err := l.Remove(); err != nil {
		return err
	}
	return l.setup()
}

func parseBaseOffset(filename string) (offset uint64, isValid bool, err error) {
	// Parse the base offset from the file name (store/index)
	return // your implementation here
}
