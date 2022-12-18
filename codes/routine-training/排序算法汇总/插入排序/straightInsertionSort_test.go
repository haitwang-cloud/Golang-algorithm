package main

import "testing"

func BenchmarkStraightInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StraightInsertionSort([]int{1, 4, 2, 7, 9, 6, 5, 8, 3, 10})
	}
}

func BenchmarkBubbleSortSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StraightInsertionSort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}
