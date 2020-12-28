package boltdb

import (
	"search-engine/chapter7/store"
	bolt "go.etcd.io/bbolt"
)

type Reader struct {
	store  *Store
	tx     *bolt.Tx
	bucket *bolt.Bucket
}

func (r *Reader) Get(key []byte) ([]byte, error) {
	var rv []byte
	v := r.bucket.Get(key)
	if v != nil {
		rv = make([]byte, len(v))
		copy(rv, v)
	}
	return rv, nil
}

func (r *Reader) PrefixIterator(prefix []byte) store.KVIterator {
	cursor := r.bucket.Cursor()

	rv := &Iterator{
		cursor: cursor,
		prefix: prefix,
	}

	rv.Seek(prefix)
	return rv
}

func (r *Reader) Close() error {
	return r.tx.Rollback()
}
