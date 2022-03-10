package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1, index2 := m-1, n-1
	indexMerge := m + n - 1
	for index2 >= 0 {
		if index1 < 0 || nums1[index1] < nums2[index2] {
			nums1[indexMerge] = nums2[index2]
			indexMerge--
			index2--

		} else {
			nums1[indexMerge] = nums1[index1]
			indexMerge--
			index1--
		}
	}
}

func main() {
	nums1, nums2, m, n := []int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}, 3, 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
