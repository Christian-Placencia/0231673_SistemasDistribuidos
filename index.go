package log

import (
	"encoding/binary"
	"io"
	"os"
	"sync"

	"github.com/tysonmote/gommap"
)

var (
	offWidth uint64 = 4                   // size of the offset field (uint32)
	posWidth uint64 = 8                   // size of the position field (uint64)
	entWidth        = offWidth + posWidth // total size of an entry
)

type index struct {
	file *os.File
	mmap gommap.MMap
	size uint64
	mu   sync.Mutex
}

func (i *index) Name() string {
	return i.file.Name()
}

// newIndex creates a new index, memory maps the file, and prepares the size.
func newIndex(f *os.File, c Config) (*index, error) {
	idx := &index{
		file: f,
	}

	fi, err := os.Stat(f.Name())
	if err != nil {
		return nil, err
	}
	idx.size = uint64(fi.Size())

	if err := os.Truncate(f.Name(), int64(c.Segment.MaxIndexBytes)); err != nil {
		return nil, err
	}

	// Memory-map the file
	if idx.mmap, err = gommap.Map(
		idx.file.Fd(),
		gommap.PROT_READ|gommap.PROT_WRITE,
		gommap.MAP_SHARED,
	); err != nil {
		return nil, err
	}

	return idx, nil
}

// Read returns the position of the given offset.
// If offset is -1, it returns the last offset and its position.
func (i *index) Read(in int64) (out uint32, pos uint64, err error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	// Check if the index is empty
	if i.size == 0 {
		return 0, 0, io.EOF
	}

	// If `in == -1`, we are asking for the last offset.
	if in == -1 {
		out = uint32((i.size / entWidth) - 1)
	} else {
		out = uint32(in)
	}

	// Calculate the offset of the entry in the mmap.
	offset := uint64(out) * entWidth
	if i.size < offset+entWidth {
		return 0, 0, io.EOF
	}

	// Read the offset and position from the memory-mapped file.
	out = binary.BigEndian.Uint32(i.mmap[offset : offset+offWidth])
	pos = binary.BigEndian.Uint64(i.mmap[offset+offWidth : offset+entWidth])
	return out, pos, nil
}

// Write adds the given offset and position to the index.
func (i *index) Write(off uint32, pos uint64) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	// Check if there's enough space in the index
	if i.size+entWidth > uint64(len(i.mmap)) {
		return io.EOF
	}

	// Write the offset and position to the memory-mapped file.
	binary.BigEndian.PutUint32(i.mmap[i.size:i.size+offWidth], off)
	binary.BigEndian.PutUint64(i.mmap[i.size+offWidth:i.size+entWidth], pos)

	// Update the size of the index.
	i.size += entWidth
	return nil
}

// Close flushes the memory-mapped file and closes the underlying file.
func (i *index) Close() error {
	i.mu.Lock()
	defer i.mu.Unlock()

	// Sync the memory-mapped file to disk.
	if err := i.mmap.Sync(gommap.MS_SYNC); err != nil {
		return err
	}

	// Truncate the file to the actual size.
	if err := os.Truncate(i.file.Name(), int64(i.size)); err != nil {
		return err
	}

	// Unmap the memory and close the file.
	if err := i.mmap.UnsafeUnmap(); err != nil {
		return err
	}
	return i.file.Close()
}
