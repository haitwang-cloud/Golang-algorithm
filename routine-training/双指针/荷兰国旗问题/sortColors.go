package main

import "fmt"

func sortColors(nums []int) {
	left, index, right := -1, 0, len(nums)
	for index < right {
		if nums[index] == 0 {
			left++
			nums[left], nums[index] = nums[index], nums[left]
			index++
		} else if nums[index] == 2 {
			right--
			nums[right], nums[index] = nums[index], nums[right]
		} else {
			index++
		}
	}
}

func swapColors(colors []int, target int) (countTarget int) {
	for index, value := range colors {
		if value == target {
			colors[index], colors[countTarget] = colors[countTarget], colors[index]
			countTarget++
		}
	}
	return
}

func sortColorsNew(nums []int) {
	index0 := swapColors(nums, 0)
	swapColors(nums[index0:], 1)
}

func main() {
	test := []int{2, 0, 1, 0, 0, 2, 1, 1, 0}
	//sortColors(test)
	sortColorsNew(test)
	fmt.Println(test)
}
