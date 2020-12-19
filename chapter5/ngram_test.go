package chapter5

import "testing"

func TestBuildNGram(t *testing.T) {

	t.Log("2-gram:")
	t.Log(Build2Gram("程咬金"))
	t.Log("n-gram:")
	t.Log(BuildNGram("程咬金", 1, 2))

}
