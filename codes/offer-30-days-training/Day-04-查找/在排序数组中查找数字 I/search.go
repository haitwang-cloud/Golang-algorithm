package main

import (
	"fmt"
)

//https://programmercarl.com/0704.二分查找.html#思路

func search(nums []int, target int) int {
	left, right := binarySearch(nums, target), binarySearch(nums, target+1)-1
	if left == len(nums) || nums[left] != target {
		return 0
	}

	return right - left + 1

}
func binarySearch(nums []int, target int) int {
	if len(nums) < 1 {
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
	fmt.Println(search(test, 6))
	fmt.Println(search(test, 8))
}
