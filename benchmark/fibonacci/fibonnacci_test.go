package fibonacci_test

import (
	"testing"

	"github.com/Chroq/benchmark/fibonacci"
)

func BenchmarkFibonnacciReccursive10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci.FibonacciRecursive(10)
	}
}

func BenchmarkFibonacciLoop10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci.FibonacciLoop(10)
	}
}

func BenchmarkFibonacciDynamic10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci.FibonacciDynamic(10)
	}
}

/*
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
*/
