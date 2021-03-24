package main

import "fmt"

//由于数组长度为n，并且要求数组内的数字都在0~n-1范围内，那么如果对数组排序的话，数组的下标可以唯一表示这个数字。
//由这个为切入点，对数组进行循环，若下标与数字不对应，则将该数字移动到对应的下标，直到下标能对应。
//若数组循环完还不能找到，则没有重复数字；若移到对应下标时与原本在该下标下的数相等，则有重复数字。
//时间复杂度O(n)，空间复杂度O(1)

func findRepeatNumber(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			} else {
				index := nums[i]
				nums[index], nums[i] = nums[i], nums[index]
			}

		}
	}
	return -1
}

func main() {
	nums1 := []int{4, 9, 5, 4, 4}
	fmt.Println(findRepeatNumber(nums1))
}
