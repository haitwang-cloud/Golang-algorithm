package main

import "fmt"

// https://leetcode.cn/problems/sort-colors/solution/yan-se-fen-lei-by-leetcode-solution/
func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left < right {
		pivot := Parition(nums, left, right)
		quickSort(nums, left, pivot-1)
		quickSort(nums, pivot+1, right)
	}
}

func Parition(nums []int, left, right int) int {
	mid := left + (right-left)/2
	if nums[left] > nums[right] {
		nums[left], nums[right] = nums[right], nums[left]
	} else if nums[mid] > nums[left] {
		nums[mid], nums[left] = nums[left], nums[mid]
	} else if nums[mid] > nums[right] {
		nums[mid], nums[right] = nums[right], nums[mid]
	}
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot

	return left

}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
