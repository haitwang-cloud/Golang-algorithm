package main

import "fmt"

//https: //leetcode.cn/problems/search-insert-position/

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	nums, test := []int{1, 3, 5, 6}, 5
	fmt.Println(searchInsert(nums, test))
}
