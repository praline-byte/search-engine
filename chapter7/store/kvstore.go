// @Description  
package store

// KVStore 是对 kv 型存储引擎的抽象.
// 提供 读，写 操作
type KVStore interface {
	Writer() (KVWriter, error)

	Reader() (KVReader, error)

	Close() error
}

// KVReader 读操作处理方法
// KVReader 被打开之后发生的写操作，不会对读操作有影响，读操作结束后应该立即关闭，以便下次读取时，访问到最新写操作之后的数据
type KVReader interface {

	// 根据 key 获取 value，如果 value 不存在则返回 nil.
	Get(key []byte) ([]byte, error)

	// 根据 prefix 获取可以访问所有符合 prefix 前缀匹配条件的 KVIterator
	PrefixIterator(prefix []byte) KVIterator

	// 关闭 KVReader
	Close() error
}

// KVIterator 定义迭代器用于遍历数据，往往用于区间查询，范围查询
// 迭代器模式，数据存储和遍历分离，为不同聚合结构提供一个统一的接口
type KVIterator interface {

	// 根据 key 将迭代器指针定位到相应位置
	// 不同存储引擎对 seek 的实现有细微差异，在 boltdb 中如果 key 不存在，会返回 key 右近邻（比指定 key 大的下一个key）
	// 如果右近邻不存在，即无数据时返回 nil
	// 我们不需要满足所有存储引擎的处理细节，只需要不同的存储引擎的实现告诉我们，当前 key 是否可用即可，即提供 Valid() bool.
	// 而是否可用属于存储引擎层面需要完成的事情，这样在当前层面上就提供了抽象，隔离了复杂，实现了通用性
	Seek(key []byte)

	// 标识迭代器当前指针是否可用(正确位置，范围外，或者无数据)
	Valid() bool

	// 将迭代器指针 +1，即指向下一个位置
	// 同样的，每当有迭代器指针移动时，我们都需要更新下 Valid 标示移动后的数据是否可用
	// 因此，实现 Next 方法也需要更新 Valid
	Next()

	// 返回迭代器当前指针指向的 key 值
	// 使用之前需要判断 Valid 状态
	Key() []byte

	// 返回迭代器当前指针指向的 value 值
	// 使用之前需要判断 Valid 状态
	Value() []byte

	// 返回当前指针的状态，指向的 key，value 的值
	Current() ([]byte, []byte, bool)

	// 关闭迭代器
	Close() error
}

// KVWriter 写操作处理方法
// 不对并发写做限制，并发写时需要注意数据安全
type KVWriter interface {

	// 写入操作
	Write(key, val []byte) error

	// 关闭 KVWriter
	Close() error
}
