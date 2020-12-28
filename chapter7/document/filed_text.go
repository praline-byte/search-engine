package document

import (
	"fmt"
	"search-engine/chapter7/analysis"
)

type TextField struct {
	name  string
	value []byte

	analyzer *analysis.Analyzer
}

func (t *TextField) Name() string {
	return t.name
}

func (t *TextField) Value() []byte {
	return t.value
}

func NewTextField(name string, value []byte) *TextField {
	return &TextField{
		name:  name,
		value: value,
	}
}

func NewTextFieldWithAnalyzer(name string, value []byte, analyzer *analysis.Analyzer) *TextField {
	return &TextField{
		name:     name,
		value:    value,
		analyzer: analyzer,
	}
}


func (t *TextField) Analyze() (int, analysis.TokenFrequencies) {
	var tokens analysis.TokenStream
	// 进行 analyze 操作
	if t.analyzer != nil {
		bytesToAnalyze := t.Value()
		tokens = t.analyzer.Analyze(bytesToAnalyze)
	} else {
		tokens = analysis.TokenStream{
			&analysis.Token{
				Start:    0,
				End:      len(t.value),
				Term:     t.value,
				Position: 1,
			},
		}
	}
	fieldLength := len(tokens)
	tokenFrequencies := analysis.TokenFrequency(tokens)
	return fieldLength, tokenFrequencies
}

func (t *TextField) String() string {
	return fmt.Sprintf("&document.TextField{Name:%s, Value: %s}", t.name, t.value)
}
