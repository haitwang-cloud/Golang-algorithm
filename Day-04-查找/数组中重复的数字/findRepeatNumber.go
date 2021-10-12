package main

import "fmt"

func findRepeatNumber(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			} else {
				index := nums[i]
				nums[i], nums[index] = nums[index], nums[i]
			}

		}
	}
	return -1
}

func main() {
	test := []int{3, 4, 1, 2, 4}
	fmt.Println(findRepeatNumber(test))
}
