package chapter5

import (
	"bytes"
	"unicode/utf8"
)

// 构建 2-gram
func Build2Gram(input string) []string {
	rv := make([]string, 0, len(input))
	runeCount := utf8.RuneCount([]byte(input))
	runes := bytes.Runes([]byte(input))
	ngramSize := 2
	for i := 0; i < runeCount; i++ {
		if i+ngramSize <= runeCount {
			rv = append(rv, string(runes[i : i+ngramSize]))
		}
	}
	return rv
}

// 构建 n-gram
func BuildNGram(input string, minLen, maxLen int) []string {
	rv := make([]string, 0, len(input))
	runeCount := utf8.RuneCount([]byte(input))
	runes := bytes.Runes([]byte(input))
	for i := 0; i < runeCount; i++ {
		for ngramSize := minLen; ngramSize <= maxLen; ngramSize++ {
			if i+ngramSize <= runeCount {
				rv = append(rv, string(runes[i : i+ngramSize]))
			}
		}
	}
	return rv
}
