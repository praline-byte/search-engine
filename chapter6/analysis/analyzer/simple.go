package analyzer

import (
	"search-engine/chapter6/analysis"
	"search-engine/chapter6/analysis/token"
	"search-engine/chapter6/analysis/tokenizer"
)

// 定义简单分析器 = 不分词 + N-Gram
func SimpleAnalyzer() (*analysis.Analyzer, error) {
	analyzer := &analysis.Analyzer{
		Tokenizer: tokenizer.NewSingleTokenTokenizer(),
		TokenFilters: []analysis.TokenFilter{
			token.NewNgramFilter(1, 2),
		},
	}
	return analyzer, nil
}