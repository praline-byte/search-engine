// @Description  
package chapter4

import "testing"

func TestSearch(t *testing.T) {

	// 初始化 引擎
	engine := NewSearchEngine("./tmp_db", "demo")

	// 准备入库索引的数据
	data := []string{
		"程咬金","孙尚香","安琪拉","程龙",
	}

	// 建立索引
	for _, name := range data {
		if err := engine.Index(name); err == nil {
			t.Logf("索引了数据:【%v】", name)
		}
	}
	t.Log("-------")

	// 进行检索
	query := "程咬金"
	t.Logf("Get 检索:【%v】", query)
	resp, _ := engine.TermSearch(query)
	t.Logf("检索结果:【%v】", resp)
	t.Log("-------")

	query = "程"
	t.Logf("前缀 检索:【%v】", query)
	resps, _ := engine.PrefixSearch(query)
	for _, bytes := range resps {
		t.Logf("检索结果:【%v】",  bytes)
	}

}
