package main

import "fmt"

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	pivot := nums[0]
	left, right := []int{}, []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i] < pivot {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}

func main() {
	nums := []int{1, 4, 2, 7, 9, 6, 5, 8, 3, 10}
	fmt.Println(quickSort(nums))
}
