package log

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

var (
	enc = binary.BigEndian
)

const (
	lenWidth = 8
)

type store struct {
	*os.File
	mu   sync.Mutex
	buf  *bufio.Writer
	size uint64
}

// newStore initializes a store with a buffered writer and sets the current size of the store.
func newStore(f *os.File) (*store, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	return &store{
		File: f,
		buf:  bufio.NewWriter(f),
		size: uint64(fi.Size()),
	}, nil
}

// Append writes the given byte slice to the store and returns the number of bytes written and its position.
func (s *store) Append(p []byte) (n uint64, pos uint64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	pos = s.size

	// Write the length of the data
	if err := binary.Write(s.buf, enc, uint64(len(p))); err != nil {
		return 0, 0, err
	}

	// Write the actual data
	w, err := s.buf.Write(p)
	if err != nil {
		return 0, 0, err
	}

	// Update the size of the store
	w += lenWidth
	s.size += uint64(w)

	return uint64(w), pos, nil
}

// Read reads the data stored at the given position.
func (s *store) Read(pos uint64) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Flush the buffer before reading
	if err := s.buf.Flush(); err != nil {
		return nil, err
	}

	// Read the length of the data
	sizeBuf := make([]byte, lenWidth)
	if _, err := s.File.ReadAt(sizeBuf, int64(pos)); err != nil {
		return nil, err
	}
	size := enc.Uint64(sizeBuf)

	// Read the actual data
	data := make([]byte, size)
	if _, err := s.File.ReadAt(data, int64(pos+lenWidth)); err != nil {
		return nil, err
	}

	return data, nil
}

// Close flushes the buffer and closes the file.
func (s *store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Flush the buffer
	if err := s.buf.Flush(); err != nil {
		return err
	}

	// Close the file
	return s.File.Close()
}

// Read implements the io.Reader interface for store.
func (s *store) Read(p []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.File.Read(p)
}
