package main

import "fmt"

// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

func findRepeatNumber(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	for index := 0; index < len(nums); index++ {
		for index != nums[index] {
			if nums[index] == nums[nums[index]] {
				return nums[index]
			}
			valueIndex := nums[index]
			nums[index], nums[valueIndex] = nums[valueIndex], nums[index]
		}
	}
	return -1
}

func main() {
	test := []int{3, 4, 2, 1, 1, 0}
	fmt.Println(findRepeatNumber(test))
}
