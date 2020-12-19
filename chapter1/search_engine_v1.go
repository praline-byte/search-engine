package chapter1


// 声明存储结构
var	store = make([]string, 0)

// 索引数据
func Index(data string) {
	store = append(store, data)
}

// 检索数据
func Search(query string) string {
	for _, name := range store {
		if name == query {
			return name
		}
	}
	return ""
}