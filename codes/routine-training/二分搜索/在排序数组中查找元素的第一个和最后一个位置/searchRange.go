package main

//https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	leftValue := binarySearch(nums, target)
	if leftValue == len(nums) || nums[leftValue] != target {
		return []int{-1, -1}
	}
	rightValue := binarySearch(nums, target+1) - 1
	return []int{leftValue, rightValue}

}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left

}
