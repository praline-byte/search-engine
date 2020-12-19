package chapter3

import "testing"

func TestSearch1(t *testing.T) {

	// 初始化 db
	config := make(map[string]interface{})
	config["path"] = "./tmp_db"
	db, err := New(config)
	if err != nil {
		t.Fatal(err)
	}

	// 准备入库索引的数据
	data := []string{
		"程咬金","孙尚香","安琪拉","程龙",
	}

	// 建立索引
	for _, name := range data {
		// 替换掉 ch1 中使用的 Index 方法，改为本次的 db.Put
		// Index(name)
		if err := db.Put([]byte(name), []byte(name)); err == nil {
			t.Logf("索引了数据:【%v】", name)
		}
	}
	t.Log("-------")

	// 进行检索
	query := "程咬金"
	t.Logf("Get 检索:【%v】", query)
	// 替换掉 ch1 中使用的 Search 方法，改为本次的 db.Get
	// resp := Search(query)
	resp, _ := db.Get([]byte(query))
	t.Logf("检索结果:【%v】", string(resp))
	t.Log("-------")

	query = "程"
	t.Logf("前缀 检索:【%v】", query)
	// 替换掉 ch1 中使用的 Search 方法，改为本次的 db.PrefixGet
	// resp := Search(query)
	resps, _ := db.PrefixGet([]byte(query))
	for _, bytes := range resps {
		t.Logf("检索结果:【%v】",  string(bytes))
	}
}
