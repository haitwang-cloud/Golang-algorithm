package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	leftValue := binarySearch(nums, target, true)
	if leftValue == len(nums) || nums[leftValue] != target {
		return []int{-1, -1}
	}
	rightValue := binarySearch(nums, target, false) - 1
	return []int{leftValue, rightValue}

}

func binarySearch(nums []int, target int, lower bool) int {
	left, right, ans := 0, len(nums)-1, len(nums)
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target || (lower && nums[mid] >= target) {
			right = mid - 1
			ans = mid

		} else {
			left = mid + 1

		}
	}

	return ans

}

func main() {
	nums, test := []int{2, 2}, 2
	fmt.Println(searchRange(nums, test))
}
