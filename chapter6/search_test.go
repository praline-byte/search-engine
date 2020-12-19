package chapter6

import (
	"search-engine/chapter6/analysis/analyzer"
	"testing"
)

// 解决复杂问题最行之有效的办法往往就是 分解，
func TestToken(t *testing.T) {
	myAnalyzer, _ := analyzer.SimpleAnalyzer()

	tokenStream := myAnalyzer.Analyze([]byte("程咬金"))
	for _, token := range tokenStream {
		t.Log(token.String())
	}
}


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
	// 这时就可以高效完成部分匹配了
	query := "咬金"
	t.Logf("Get 检索:【%v】", query)
	resp, _ := engine.TermSearch(query)
	t.Logf("检索结果:【%v】", resp)
	t.Log("-------")

}
