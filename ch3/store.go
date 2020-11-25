package ch3

import (
	"bytes"
	"fmt"
	"os"

	bolt "go.etcd.io/bbolt"
)


type Store struct {
	path   string   // db 文件存储路径
	bucket string   // bucket 存放 kv 的大池子，类似 table
	db     *bolt.DB // boltDb 实例
}

// 根据 config 配置参数初始化 store
func New(config map[string]interface{}) (*Store, error) {

	// 通过配置获取存储路径和 B+树的名称
	path, ok := config["path"].(string)
	if !ok {
		return nil, fmt.Errorf("must specify path")
	}
	if path == "" {
		return nil, os.ErrInvalid
	}

	bucket, ok := config["bucket"].(string)
	if !ok {
		bucket = "liangzai"
	}

	// 可以通过 option 定制 bolt，我们这里使用原始的就可以
	bo := &bolt.Options{}

	// 在指定位置上，创建数据库文件，并赋予读写权限
	db, err := bolt.Open(path, 0600, bo)
	if err != nil {
		return nil, err
	}

	// 创建 B+ 树
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		return err
	})

	// 返回我们的存储结构
	rv := Store{
		path:        path,
		bucket:      bucket,
		db:          db,
	}
	return &rv, nil
}

// boltDB 规定只读事务结束时，必须要调用 Rollback，以此来释放脏页，保证事务的隔离性
// 关于事务这里可以先简单理解为"一件有头有尾的完整的事情"，我们会在之后更详细的了解事务，事务的特性，事务的隔离级别和不同级别的实现标准等等。

// 根据 key 获取 b+树中的 val
func (bs *Store) Get(key []byte) ([]byte, error) {

	// 开启一个只读事务，这个事务不会产生任何更新，指定 writable 为 false
	tx, err := bs.db.Begin(false)

	// boltDB 规定只读事务结束时，必须要调用 Rollback，以此来释放脏页，保证事务的隔离特性
	defer tx.Rollback()

	if err != nil {
		return nil, err
	}

	// 获取指定的 b+树
	bucket := tx.Bucket([]byte(bs.bucket))

	// 获取 b+树中的 key-value
	return bucket.Get(key), nil
}

// 将 key-val 数据插入到 b+树中
func (bs *Store) Put(key, val []byte) (err error) {

	// 开启一个只读事务，这个事务将会修改数据，指定 writable 为 true
	tx, err := bs.db.Begin(true)

	if err != nil {
		return nil
	}

	// 如果此次 Put 有失败，则回滚事务 rollback，如果 Put 成功，则提交事务 commit
	defer func() {
		if err == nil {
			err = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}()

	// 获取指定的 b+树
	bucket := tx.Bucket([]byte(bs.bucket))

	// 向 b+树中插入数据
	err = bucket.Put(key, val)
	return err
}

// 关闭资源
func (bs *Store) Close() error {
	return bs.db.Close()
}


// 前缀检索
func (bs *Store) PrefixGet(prefix []byte) ([][]byte, error) {

	// 声明存放返回结果的容器
	vals := make([][]byte, 0)

	// 开启一个读事务，与 Get() 方法一致
	tx, err := bs.db.Begin(false)
	defer tx.Rollback()

	if err != nil {
		return nil, err
	}

	// 获取指定的 b+树
	bucket := tx.Bucket([]byte(bs.bucket))

	// 获取 b+树中用于遍历的指针
	cursor := bucket.Cursor()

	// 这是一个复合结构
	// k, v := cursor.Seek(prefix) 表示执行这段循环时，初始化要做的事情，这里我们把要匹配的前缀，通过游标的方式定位到 b+中
	// bytes.HasPrefix(k, prefix) 判断游标在当前位置时，是否满足 "前缀"条件
	// k, v = cursor.Next() 如果满足条件，就将指针向下移动一位
	// 因为在 b+树中所有数据有序排列，所以不断移动指针并判断是否满足条件，就可以找到所有目标结果，如果遇到不满足的结果，就会退出当前循环
	for k, v := cursor.Seek(prefix); bytes.HasPrefix(k, prefix); k, v = cursor.Next() {
		vals = append(vals, v)
	}
	return vals, nil
}

// db 只有一个实例，只加载一次
// reader 和 writer 每次都要开启事务，可以拆分，按需加载

// 存在两个问题
// 1. 存在大量重复代码，Get(), MultiGet(), PrefixSearch(), RangeSearch() 都存在重复的`开启只读事务`,`调用 db 原生方法获取数据`,`将结果数据返回`，
// `关闭只读事务`，当重复代码出现我们就要嗅到改造它的时机到了（当然要考虑业务阶段，改造成本，团队规范和心情）
// 2. 如果我们业务升级，当前的 db 引擎不再满足需求时（别忘了 boltDb 的主要特点就是简单，易用，可靠，读场景多于写场景），改用这些底层 API 意味着大量的调用方改动成本
// 也随之暴涨。
// 如何隔离复杂，解耦依赖，拥抱变化？





