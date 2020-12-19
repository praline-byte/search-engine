package token

import (
	"bytes"
	"search-engine/chapter6/analysis"
	"unicode/utf8"
)

type NgramFilter struct {
	minLength int
	maxLength int
}

func NewNgramFilter(minLength, maxLength int) analysis.TokenFilter {
	return &NgramFilter{
		minLength: minLength,
		maxLength: maxLength,
	}
}

func (s *NgramFilter) Filter(input analysis.TokenStream) analysis.TokenStream {
	rv := make(analysis.TokenStream, 0, len(input))

	for _, token := range input {
		runeCount := utf8.RuneCount(token.Term)
		runes := bytes.Runes(token.Term)
		for i := 0; i < runeCount; i++ {
			for ngramSize := s.minLength; ngramSize <= s.maxLength; ngramSize++ {
				if i+ngramSize <= runeCount {
					ngramTerm := analysis.BuildTermFromRunes(runes[i : i+ngramSize])
					token := analysis.Token{
						Position: token.Position,
						Start:    token.Start,
						End:      token.End,
						Term:     ngramTerm,
					}
					rv = append(rv, &token)
				}
			}
		}
	}

	return rv
}

