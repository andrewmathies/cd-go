package main

import (
	"testing"
)

func BenchmarkCompare(b *testing.B) {
	max := 5
	n := 100000000
	k := 2

	arr := generateArr(n, max)

	b.Run("naive", func(b *testing.B) {
		naiveSum(arr, n, max)
	})

	b.Run("divideAndConquer", func(b *testing.B) {
		divideAndConquerSum(arr, n, k, max)
	})
}
