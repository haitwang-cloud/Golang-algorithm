package main

import (
	"fmt"
	"math/rand"
	"time"
)

//快排可以解决问题，但是它需要确定数组中所有元素的正确位置，
//对于本题而言，我们只需要确定第k大元素的位置pos,我们只需要确保pos左边的元素都比它小，
//pos右边的元素都比它大即可，不需要关心其左边和右边的集合是否有序，
//所以，我们需要对快排进行改进(以下称分区)，将目标值的位置pos与分区函数Partition求得的位置index进行比对，
//如果两值相等，说明index对应的元素即为所求值，如果index<pos，则递归的在[index+1, right]范围求解；
//否则则在[left, index-1]范围求解，如此便可大幅缩小求解范围。

func findKthLargest(nums []int, k int) int {
	TopKSplit(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func Parition(nums []int, start, stop int) int {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int()%(stop-start+1) + start
	nums[picked], nums[start] = nums[start], nums[picked]

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

//获得前k小的数
//
//
//func TopkSmallest(nums []int, k int) []int {
//	TopkSplit(nums, k, 0, len(nums)-1)
//	return nums[:k]
//}
//2 获得前k大的数
//
//
//func TopkLargest(nums []int, k int) []int {
//	TopkSplit(nums, len(nums)-k, 0, len(nums)-1)
//	return nums[len(nums)-k:]
//}
//3 获取第k小的数
//
//
//func TopkSmallestElement(nums []int, k int) int {
//	TopkSplit(nums, k, 0, len(nums)-1)
//	return nums[k-1]
//}
//4 获取第k大的数
//
//
//func TopkLargestElement(nums []int, k int) int {
//	TopkSplit(nums, len(nums)-k, 0, len(nums)-1)
//	return nums[len(nums)-k]
//}

func main() {
	nums1, k := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 20
	fmt.Println(findKthLargest(nums1, k))
}
