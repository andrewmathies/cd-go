package main

import (
	"math/rand"
	"time"
)

// var max, n, k int

func generateArr(n int, max int) []uint64 {
	rand.Seed(time.Now().UnixNano())
	arr := make([]uint64, n)

	for i := 0; i < n; i++ {
		arr[i] = uint64(rand.Intn(max))
	}

	return arr
}

func chunk(arr []uint64, ch chan uint64) {
	total := uint64(0)

	for _, val := range arr {
		total += val
	}

	ch <- total
}

func divideAndConquerSum(arr []uint64, n int, k int, max int) uint64 {
	total := uint64(0)
	ch := make(chan uint64)

	for i := 0; i < k; i++ {
		top := (i + 1) * (n / k)
		bottom := i * (n / k)
		go chunk(arr[bottom:top], ch)
	}

	for i := 0; i < k; i++ {
		total += <-ch
	}

	return total
}

func naiveSum(arr []uint64, n int, max int) uint64 {
	total := uint64(0)

	for _, val := range arr {
		total += val
	}

	return total
}

func main() {
	/*
		flag.IntVar(&max, "max", 10, "max number an element in array can be")
		flag.IntVar(&n, "length", 1000, "number of elements in array")
		flag.IntVar(&k, "threads", 10, "number of threads divide and conquer strat will use")
		flag.Parse()
	*/
}
