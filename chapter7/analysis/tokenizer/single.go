// @Description  
// @Author  wupeng.0
package tokenizer

import (
	"search-engine/chapter7/analysis"
)

// 不需要分词的 Tokenizer
type singleTokenTokenizer struct {
}

func NewSingleTokenTokenizer() *singleTokenTokenizer {
	return &singleTokenTokenizer{}
}

func (t *singleTokenTokenizer) Tokenize(input []byte) analysis.TokenStream {
	return analysis.TokenStream{
		&analysis.Token{
			Term:     input,
			Position: 1,
			Start:    0,
			End:      len(input),
		},
	}
}