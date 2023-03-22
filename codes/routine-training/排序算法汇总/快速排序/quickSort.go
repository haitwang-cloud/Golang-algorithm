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
	} else if nums[mid] > nums[left] {
		nums[mid], nums[left] = nums[left], nums[mid]
	} else if nums[mid] > nums[right] {
		nums[right], nums[mid] = nums[mid], nums[right]
	}
	pivot := nums[left]
	for left < right {
		/*
			从右往左遍历,判断 nums[right] 是否大于基准数，如果大于则左移 right 指针，
			直至找到一个小于基准数的元素，将其填入之前的坑中
		*/
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		/*
			从左往右遍历,判断 nums[left] 是否小于基准数，如果小于则左移 left 指针，
			直至找到一个大于基准数的元素，将其填入之前的坑中
		*/
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
