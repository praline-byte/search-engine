package boltdb

import (
	"fmt"
	"os"
	"search-engine/ch4/store"

	bolt "go.etcd.io/bbolt"
)

type Store struct {
	path   string
	bucket string
	db     *bolt.DB
}

func New(config map[string]interface{}) (store.KVStore, error) {
	path, ok := config["path"].(string)
	if !ok {
		return nil, fmt.Errorf("must specify path")
	}
	if path == "" {
		return nil, os.ErrInvalid
	}

	bucket, ok := config["bucket"].(string)
	if !ok {
		bucket = "demo"
	}

	bo := &bolt.Options{}
	db, err := bolt.Open(path, 0600, bo)
	if err != nil {
		return nil, err
	}

	// 创建 B+ 树
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})

	rv := Store{
		path:   path,
		bucket: bucket,
		db:     db,
	}
	return &rv, nil
}

func (bs *Store) Reader() (store.KVReader, error) {
	tx, err := bs.db.Begin(false)
	if err != nil {
		return nil, err
	}
	return &Reader{
		store:  bs,
		tx:     tx,
		bucket: tx.Bucket([]byte(bs.bucket)),
	}, nil
}

func (bs *Store) Writer() (store.KVWriter, error) {
	return &Writer{
		store: bs,
	}, nil
}

func (bs *Store) Close() error {
	return bs.db.Close()
}
