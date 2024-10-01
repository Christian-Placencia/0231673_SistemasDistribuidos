package log

type storeReader struct {
	store *store
	pos   uint64
}

// Implement the Read method for storeReader
func (sr *storeReader) Read(p []byte) (int, error) {
	data, err := sr.store.Read(sr.pos)
	if err != nil {
		return 0, err
	}

	n := copy(p, data)
	sr.pos += uint64(n)
	return n, nil
}
