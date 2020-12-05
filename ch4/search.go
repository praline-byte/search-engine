package ch4

import (
	"search-engine/ch4/store"
	"search-engine/ch4/store/boltdb"
)

// 升级之后的搜索引擎，我们只要持有 KVStore 接口定义就可以了，KVStore 的实现交给初始化去完成
type searchEngine struct {
	store store.KVStore
}

// storeName 索引文件的存储目录
func NewSearchEngine(path, bucket string) *searchEngine {

	// 定义我们 store 层的配置
	// 这里我们制定 boltdb 加载路径和 bucket
	config := make(map[string]interface{})
	config["path"] = path
	config["bucket"] = bucket

	// 调用一个 KVStore 的实现，这里是 boltdb
	store, err := boltdb.New(config)
	if err != nil {
		panic(err)
	}
	return &searchEngine{
		store: store,
	}
}

// 索引数据
func (engine *searchEngine) Index(data string) (err error) {
	if writer, err := engine.store.Writer(); err == nil {
		defer writer.Close()
		return writer.Write([]byte(data), []byte(data))
	}
	return err
}

// 完全匹配查询
func (engine *searchEngine) TermSearch(query string) (found string, err error) {
	if reader, err := engine.store.Reader(); err == nil {
		defer reader.Close()

		if bytes, err := reader.Get([]byte(query)); err == nil {
			return string(bytes), nil
		}
	}
	return
}

// 前缀匹配查询
func (engine *searchEngine) PrefixSearch(query string) (found []string, err error) {
	if reader, err := engine.store.Reader(); err == nil {
		defer reader.Close()

		// 迭代器模式，定位游标初始位置
		iter := reader.PrefixIterator([]byte(query))

		// 判断是否游标可用
		_, val, valid := iter.Current()
		for valid {
			// 满足匹配条件，放入结果集并移动游标
			found = append(found, string(val))
			iter.Next()

			// 获取最新游标指向的值并更新状态，进行下次循环
			_, val, valid = iter.Current()
		}
	}
	return
}

// 关闭存储引擎
func (engine *searchEngine) Close() error {
	return engine.store.Close()
}
