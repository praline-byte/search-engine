package document

import (
	"fmt"
	"search-engine/chapter7/analysis"
)

type BooleanField struct {
	name  string
	value []byte
}

func (b *BooleanField) Name() string {
	return b.name
}

func (b *BooleanField) Analyze() (int, analysis.TokenFrequencies) {
	tokens := make(analysis.TokenStream, 0)
	tokens = append(tokens, &analysis.Token{
		Start:    0,
		End:      len(b.value),
		Term:     b.value,
		Position: 1,
	})

	fieldLength := len(tokens)
	tokenFrequencies := analysis.TokenFrequency(tokens)
	return fieldLength, tokenFrequencies
}

func (b *BooleanField) Value() []byte {
	return b.value
}

func (b *BooleanField) Boolean() bool {
	return bytesToBools(b.value)[0]
}

func (b *BooleanField) String() string {
	return fmt.Sprintf("&document.BooleanField{Name:%s, Value: %s}", b.name, b.value)
}

func NewBooleanFieldFromBytes(name string, value []byte) *BooleanField {
	return &BooleanField{
		name:  name,
		value: value,
	}
}

func NewBooleanField(name string, b bool) *BooleanField {
	return NewBooleanFieldFromBytes(name, boolsToBytes([]bool{b}))
}

func boolsToBytes(t []bool) []byte {
	b := make([]byte, (len(t)+7)/8)
	for i, x := range t {
		if x {
			b[i/8] |= 0x80 >> uint(i%8)
		}
	}
	return b
}

func bytesToBools(b []byte) []bool {
	t := make([]bool, 8*len(b))
	for i, x := range b {
		for j := 0; j < 8; j++ {
			if (x<<uint(j))&0x80 == 0x80 {
				t[8*i+j] = true
			}
		}
	}
	return t
}
