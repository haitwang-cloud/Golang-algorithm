package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid

		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	nums, test := []int{1, 3, 5, 6}, 7
	fmt.Println(searchInsert(nums, test))
}
