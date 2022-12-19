package main

import "testing"

func BenchmarkShellSortSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShellSort([]int{1, 4, 2, 7, 9, 6, 5, 8, 3, 10})
	}
}

func BenchmarkShellSortSorted(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShellSort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}
