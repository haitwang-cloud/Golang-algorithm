package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

func main() {
	test, k := []int{2, 3, 1, 5, 6}, 3
	fmt.Println(removeElement(test, k))
}
