package main

import "fmt"

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return left
}

func main() {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println(missingNumber(test))
}
