package document

import (
	"encoding/binary"
	"fmt"
	"search-engine/chapter7/analysis"
	"math"
)

type NumericField struct {
	name  string
	value []byte
}

func (n *NumericField) Name() string {
	return n.name
}

func (n *NumericField) Analyze() (int, analysis.TokenFrequencies) {
	tokens := make(analysis.TokenStream, 0)
	tokens = append(tokens, &analysis.Token{
		Start:    0,
		End:      len(n.value),
		Term:     n.value,
		Position: 1,
	})

	fieldLength := len(tokens)
	tokenFrequencies := analysis.TokenFrequency(tokens)
	return fieldLength, tokenFrequencies
}

func (n *NumericField) Value() []byte {
	return n.value
}

func (n *NumericField) Number() float64 {
	return ByteToFloat64(n.value)
}

func (n *NumericField) String() string {
	return fmt.Sprintf("&document.NumericField{Name:%s, Value: %s}", n.name, n.value)
}

func NewNumericFieldFromBytes(name string, value []byte) *NumericField {
	return &NumericField{
		name:  name,
		value: value,
	}
}

func NewNumericField(name string, number float64) *NumericField {
	return NewNumericFieldFromBytes(name, Float64ToByte(number))
}

//Float64ToByte Float64转byte
func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

//ByteToFloat64 byte转Float64
func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}
