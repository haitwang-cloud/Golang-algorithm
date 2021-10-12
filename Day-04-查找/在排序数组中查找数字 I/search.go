package main

import (
	"fmt"
	"sort"
)
//https://programmercarl.com/0704.二分查找.html#思路

func search(nums []int, target int) int {
	first, second := sort.SearchInts(nums, target), binarySearch(nums, target+1)-1
	if first == len(nums) || nums[first] != target {
		return 0
	}
	return second - first + 1

}
func binarySearch(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return left
}
func main() {
	test := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(search(test, 8))
	fmt.Println(search(test, 6))
}
