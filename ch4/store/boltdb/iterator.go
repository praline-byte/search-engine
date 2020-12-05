package boltdb

import (
	"bytes"

	bolt "go.etcd.io/bbolt"
)

type Iterator struct {
	cursor *bolt.Cursor // boltDB 用于迭代遍历的游标
	prefix []byte       // 指定遍历的前缀，用于前缀搜索
	start  []byte       // 指定遍历左区间，用于区间搜索
	end    []byte       // 指定遍历右区间，用于区间搜索
	valid  bool         // 遍历是否终止
	key    []byte       // 遍历到当前位置的 key 值
	val    []byte       // 遍历到当前位置的 val 值
}

// 将游标定位到 k 的位置
// 不同的引擎实现原理不一样，以 boltDB 为例，其实现在 b+tree 上，会从根部开始二分查找
// 会有两种情况，如果找到了，返回当前的 k-v，seek 流程结束
// 如果不存在，boltDB 会返回 k 由
func (i *Iterator) Seek(k []byte) {
	if i.start != nil && bytes.Compare(k, i.start) < 0 {
		k = i.start
	}
	if i.prefix != nil && !bytes.HasPrefix(k, i.prefix) {
		if bytes.Compare(k, i.prefix) < 0 {
			k = i.prefix
		} else {
			i.valid = false
			return
		}
	}
	i.key, i.val = i.cursor.Seek(k)
	i.updateValid()
}

// 移动游标到下一个位置，并更新遍历状态是否终止
func (i *Iterator) Next() {
	i.key, i.val = i.cursor.Next()
	i.updateValid()
}

func (i *Iterator) updateValid() {
	i.valid = i.key != nil
	if i.valid {
		if i.prefix != nil {
			i.valid = bytes.HasPrefix(i.key, i.prefix)
		} else if i.end != nil {
			i.valid = bytes.Compare(i.key, i.end) < 0
		}
	}
}

func (i *Iterator) Current() ([]byte, []byte, bool) {
	return i.key, i.val, i.valid
}

func (i *Iterator) Key() []byte {
	return i.key
}

func (i *Iterator) Value() []byte {
	return i.val
}

func (i *Iterator) Valid() bool {
	return i.valid
}

// boltDB 的迭代器不需要单独关闭，统一由 tx.close() 控制
func (i *Iterator) Close() error {
	return nil
}
