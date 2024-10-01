package log

type Config struct {
	CommitLog *Log
	Segment   struct {
		MaxStoreBytes uint64
		MaxIndexBytes uint64
		InitialOffset uint64
	}
}
