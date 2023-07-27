package main

import (
	"fmt"
)

func main() {
	f := func(x int) float64 {
		return float64(x * x)
	}
	g := func(x float64) bool {
		return x < 10
	}
	h := func(x, acc float64) float64 {
		return x + acc
	}

	// [0,1,2,3,4] -> [0,1,4,9,16] -> [0,1,4,9] -> 14
	fmt.Println(Reduce(Filter(Map(iterator(5), f), g), h)) // 14
}

func iterator(n int) func(yield func(int) bool) bool {
	return func(yield func(int) bool) bool {
		for i := range n {
			if !yield(i) {
				return false
			}
		}
		return true
	}
}

func Map[T, S any](iter func(func(T) bool) bool, f func(T) S) func(yield func(S) bool) bool {
	return func(yield func(S) bool) bool {
		for x := range iter {
			if !yield(f(x)) {
				return false
			}
		}
		return true
	}
}

func Filter[T any](iter func(func(T) bool) bool, f func(T) bool) func(yield func(T) bool) bool {
	return func(yield func(T) bool) bool {
		for x := range iter {
			if !f(x) {
				continue
			}
			if !yield(x) {
				return false
			}
		}
		return true
	}
}

func Reduce[T, S any](iter func(func(T) bool) bool, f func(T, S) S) (acc S) {
	for x := range iter {
		acc = f(x, acc)
	}
	return
}
