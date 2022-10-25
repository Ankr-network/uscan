package kv

type ReadOption struct {
	table      string
	Latest     bool
	NextIfNone bool
}

type WriteOption struct {
	table string
}

type Putter interface {
	Put(key, val []byte, opts *WriteOption)
}

type Getter interface {
	Get(key []byte, opts *ReadOption) ([]byte, error)
}

type Closer interface {
	Close() error
}

type Database interface {
	Putter
	Getter
	Closer
}
