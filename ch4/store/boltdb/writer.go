package boltdb

type Writer struct {
	store *Store
}

func (w *Writer) Write(key, val []byte) (err error) {
	tx, err := w.store.db.Begin(true)
	if err != nil {
		return
	}
	// 原子性保证，要么成功 commit 提交事务, 要么失败 rollback 回滚事务
	defer func() {
		if err == nil {
			err = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}()

	bucket := tx.Bucket([]byte(w.store.bucket))
	err = bucket.Put(key, val)
	return
}

func (w *Writer) Close() error {
	return nil
}
