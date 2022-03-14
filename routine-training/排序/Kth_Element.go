package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	TopKSplit(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func Parition(nums []int, start, stop int) int {

	if start >= stop {
		return -1
	}
	pivot := nums[start]
	left, right := start, stop
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] < pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}

func TopKSplit(nums []int, k, start, stop int) {
	if start < stop {
		index := Parition(nums, start, stop)
		if index == k {
			return
		} else if index < k {
			TopKSplit(nums, k, index+1, stop)
		} else {
			TopKSplit(nums, k, start, index-1)
		}
	}
}

func findKthLargestNew(nums []int, k int) int {
	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(a []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(a, i, heapSize)
	}
}

func maxHeapify(a []int, i, heapSize int) {
	left, right, largest := i*2+1, i*2+2, i
	if left < heapSize && a[left] > a[largest] {
		largest = left
	}
	if right < heapSize && a[right] > a[largest] {
		largest = right
	}
	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		maxHeapify(a, largest, heapSize)
	}
}

func main() {
	nums1, k := []int{3, 2, 1, 5, 6, 4}, 2
	fmt.Println(findKthLargest(nums1, k))
	fmt.Println(findKthLargestNew(nums1, k))
	nums2 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println(findKthLargest(nums2, k))
	fmt.Println(findKthLargestNew(nums2, k))
}
