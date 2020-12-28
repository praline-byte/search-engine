package document

import "search-engine/chapter7/analysis"

type Field interface {
	Name() string  // 返回字段名称
	Value() []byte // 返回字段内容

	Analyze() (int, analysis.TokenFrequencies) // 对字段进行 Analyze 处理，返回 token 的个数及 term 的频率信息
}
