package main

import "fmt"

func quickSort(nums []int) []int {
	_quickSort(nums, 0, len(nums)-1)
	return nums
}

func _quickSort(nums []int, left, right int) {
	if left < right {
		pivot := partition(nums, left, right)
		_quickSort(nums, left, pivot-1)
		_quickSort(nums, pivot+1, right)
	}
}

func partition(nums []int, left, right int) int {
	//三数取中
	mid := left + (right-left)/2
	if nums[left] > nums[right] {
		nums[left], nums[right] = nums[right], nums[left]
	} else if nums[mid] > nums[right] {
		nums[mid], nums[right] = nums[right], nums[mid]
	} else if nums[mid] > nums[left] {
		nums[mid], nums[left] = nums[left], nums[mid]
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
	nums := []int{1, 4, 2, 7, 9, 6, 5, 8, 3, 10}
	fmt.Println(quickSort(nums))
}
