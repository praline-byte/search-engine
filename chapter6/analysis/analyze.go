// @Description  
// @Author  wupeng.0
package analysis

import "fmt"

// Token 是搜索引擎中的一等公民，是我们用以索引的最小单位
type Token struct {

	// 表示 Token 首字符在原字段中的开始位置
	Start int `json:"start"`

	// 表示 Token 尾字符在原字段中的结束位置
	End int `json:"end"`

	// token 的实际值
	Term []byte `json:"term"`

	// 整个 token 在原字段中的位置
	Position int `json:"position"`
}

func (t *Token) String() string {
	return fmt.Sprintf("Start: %d  End: %d  Position: %d  Token: %s ", t.Start, t.End, t.Position, string(t.Term))
}

// 一组 Token 集合，通常代表一个输入产生的 Token 结果集
type TokenStream []*Token


// Tokenizer 分词接口
// 方法 Tokenize() 表示将输入分词成一组 Token 集合
// 这一步不会破坏原字段信息，也即不会对原信息增删修改，仅用来分词，分词之后的 Token，可以无损拼接为原始字段
// 增删修改的操作，在 TokenFilter 中完成
type Tokenizer interface {
	Tokenize([]byte) TokenStream
}

// TokenFilter 过滤转换器接口
// 对一组 Token 进行增删改等操作，返回操作之后的 Token 集合
type TokenFilter interface {
	Filter(TokenStream) TokenStream
}

// Analyzer 定义分析器
// 实际上完整的分析器至少由 CharFilters + Tokenizer + TokenFilters 三部分组成
// CharFilters 用于对字符集进行转换，暂时用不到
// Tokenizer 用于分词
// TokenFilters 用于分词后 Token 的过滤转换等操作
type Analyzer struct {
	Tokenizer    Tokenizer
	TokenFilters []TokenFilter
}

func (a *Analyzer) Analyze(input []byte) TokenStream {
	tokens := a.Tokenizer.Tokenize(input)
	if a.TokenFilters != nil {
		for _, tf := range a.TokenFilters {
			tokens = tf.Filter(tokens)
		}
	}
	return tokens
}
