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
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
			if nums[i] < nums[pivot] {
					swap(nums, i, index)
					index += 1
			}
	}
	swap(nums, pivot, index-1)
	return index - 1
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	nums := []int{1, 4, 2, 7, 9, 6, 5, 8, 3, 10}
	fmt.Println(quickSort(nums))
}
