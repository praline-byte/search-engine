package chapter1

import "testing"

func TestSearchEngineV1(t *testing.T) {

	// 准备入库索引的数据
	data := []string{
		"程咬金","孙尚香","安琪拉",
	}

	// 建立索引
	for _, name := range data {
		t.Logf("索引了数据:【%v】", name)
		Index(name)
	}

	// 进行检索
	query := "程咬金"
	t.Logf("检索:【%v】", query)
	resp := Search(query)
	t.Logf("检索结果:【%v】", resp)
}