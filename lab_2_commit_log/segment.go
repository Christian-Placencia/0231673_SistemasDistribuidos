// segment.go
package log

import (
	"fmt"
	"os"
	"path"

	api "github.com/Robinthatdoesnotsuck/api/log_v1"
	"google.golang.org/protobuf/proto"
)

type Segment struct {
	store      *store
	index      *index
	baseOffset uint64
	nextOffset uint64
	config     Config
}

func newSegment(dir string, baseOffset uint64, c Config) (*Segment, error) {
	s := &Segment{
		baseOffset: baseOffset,
		config:     c,
	}

	var err error

	// Create the store file
	storeFile, err := os.OpenFile(
		path.Join(dir, fmt.Sprintf("%d.store", baseOffset)),
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	if s.store, err = newStore(storeFile); err != nil {
		return nil, err
	}

	// Create the index file
	indexFile, err := os.OpenFile(
		path.Join(dir, fmt.Sprintf("%d.index", baseOffset)),
		os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	if s.index, err = newIndex(indexFile, c); err != nil {
		return nil, err
	}

	// Set nextOffset to the last entry's offset + 1
	if off, _, err := s.index.Read(-1); err != nil {
		s.nextOffset = baseOffset
	} else {
		s.nextOffset = baseOffset + uint64(off) + 1
	}

	return s, nil
}

func (s *Segment) Append(record *api.Record) (uint64, error) {
	cur := s.nextOffset
	record.Offset = cur

	p, err := proto.Marshal(record)
	if err != nil {
		return 0, err
	}

	_, pos, err := s.store.Append(p)
	if err != nil {
		return 0, err
	}

	if err = s.index.Write(
		uint32(s.nextOffset-s.baseOffset), pos); err != nil {
		return 0, err
	}

	s.nextOffset++
	return cur, nil
}

func (s *Segment) Read(off uint64) (*api.Record, error) {
	_, pos, err := s.index.Read(int64(off - s.baseOffset))
	if err != nil {
		return nil, err
	}

	p, err := s.store.Read(pos)
	if err != nil {
		return nil, err
	}

	record := &api.Record{}
	err = proto.Unmarshal(p, record)
	return record, err
}

func (s *Segment) IsMaxed() bool {
	return s.store.size >= s.config.Segment.MaxStoreBytes || s.index.size >= s.config.Segment.MaxIndexBytes
}

func (s *Segment) Remove() error {
	if err := s.Close(); err != nil {
		return err
	}
	if err := os.Remove(s.index.Name()); err != nil {
		return err
	}
	return os.Remove(s.store.Name())
}

func (s *Segment) Close() error {
	if err := s.index.Close(); err != nil {
		return err
	}
	return s.store.Close()
}
