package main

import "fmt"

func removeElement(nums []int, val int) int {
	resultIndex := 0
	if len(nums) < 0 {
		return resultIndex
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[resultIndex] = nums[i]
			resultIndex++
		}
	}
	return resultIndex
}

func main() {
	test, k := []int{2, 3, 1, 5, 6}, 3
	fmt.Println(removeElement(test, k))
}
