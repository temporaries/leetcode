package main

import (
	"testing"
)

var s = "LVIII"

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		romanToInt(s)
	}
}

func Test(t *testing.T) {
	h := romanToInt(s)
	println(h)
}
