package analysis

// TokenFreq 表示 term 出现的次数
type TokenFreq struct {
	Term      []byte
	frequency int
}

func (tf *TokenFreq) Frequency() int {
	return tf.frequency
}

// TokenFrequencies 表示文档中所有 term 的 TokenFreq 信息
type TokenFrequencies map[string]*TokenFreq

func TokenFrequency(tokens TokenStream) TokenFrequencies {
	rv := make(map[string]*TokenFreq, len(tokens))

	for _, token := range tokens {
		curr, exists := rv[string(token.Term)]
		// 如果已经存在，则 frequency++，否则初始化为 1
		if exists {
			curr.frequency++
		} else {
			rv[string(token.Term)] = &TokenFreq{
				Term:      token.Term,
				frequency: 1,
			}
		}
	}

	return rv
}
