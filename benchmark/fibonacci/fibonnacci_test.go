package fibonacci_test

import (
	"testing"

	"github.com/Chroq/benchmark/fibonacci"
)

var result int

func BenchmarkFibonnacciReccursive10(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibonacci.FibonacciRecursive(10)
		r = fibonacci.FibonacciRecursive(10)
		r = fibonacci.FibonacciRecursive(10)
	}
	result = r
}

func BenchmarkFibonacciLoop10(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibonacci.FibonacciDynamic(10)
		r = fibonacci.FibonacciDynamic(10)
		r = fibonacci.FibonacciDynamic(10)
	}
	result = r
}

func BenchmarkFibonacciDynamic10(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibonacci.FibonacciDynamic(10)
		r = fibonacci.FibonacciDynamic(10)
		r = fibonacci.FibonacciDynamic(10)
	}
	result = r
}

func BenchmarkFibonacciMemoization10(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		memoization := fibonacci.NewMemoization(10)
		r = memoization.Get(10)
		r = memoization.Get(10)
		r = memoization.Get(10)
	}
	result = r
}
