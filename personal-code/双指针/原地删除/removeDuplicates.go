package main

import "fmt"

func removeDuplicates(nums []int) int {
	resultIndex := 0
	if len(nums) < 0 {
		return resultIndex
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[resultIndex] {
			resultIndex++
			nums[resultIndex] = nums[i]
		}
	}
	return resultIndex + 1
}

func main() {
	test := []int{1, 1, 3, 5, 6}
	fmt.Println(removeDuplicates(test))
}
