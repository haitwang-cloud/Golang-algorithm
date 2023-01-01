package main

import (
	"fmt"
)

/*
快排可以解决问题，但是它需要确定数组中所有元素的正确位置，
对于本题而言，我们只需要确定第k大元素的位置pos,我们只需要确保pos左边的元素都比它小，
pos右边的元素都比它大即可，不需要关心其左边和右边的集合是否有序，
所以，我们需要对快排进行改进(以下称分区)，将目标值的位置pos与分区函数Partition求得的位置pivot进行比对，
如果两值相等，说明pivot对应的元素即为所求值，如果pivot<pos，则递归的在[pivot+1, right]范围求解；
否则则在[left, pivot-1]范围求解，如此便可大幅缩小求解范围。
*/
func findKthLargest(nums []int, k int) int {
	TopKSplit(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func Parition(nums []int, left, right int) int {
	mid := left + (right-left)/2
	if nums[left] > nums[right] {
		nums[left], nums[right] = nums[right], nums[left]
	} else if nums[mid] > nums[right] {
		nums[right], nums[mid] = nums[mid], nums[right]
	} else if nums[mid] > nums[left] {
		nums[left], nums[mid] = nums[mid], nums[left]
	}
	pivot := nums[left]
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

func TopKSplit(nums []int, k, left, right int) {
	if left < right {
		pivot := Parition(nums, left, right)
		if pivot == k {
			return
		} else if pivot < k {
			TopKSplit(nums, k, pivot+1, right)
		} else {
			TopKSplit(nums, k, left, pivot-1)
		}
	}
}

func main() {
	nums1, k := []int{3, 2, 1, 5, 6, 4}, 2
	fmt.Println(findKthLargest(nums1, k))
}
