package main

import "fmt"

// https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	/*

		当 nums[m]>nums[right]时，m 一定在左排序数组中，即旋转点x闭区间[m+1,right]内，因此执行left=m+1；
		当 nums[m]<nums[right]时，m 一定在右排序数组中，即旋转点x闭区间[left,m]内，因此执行right=m；
		当 nums[m]==nums[right]时,无法判断m在哪个排序数组中,执行right=right−1 缩小判断范围

	*/
	left, right := 0, len(numbers)-1
	for left < right {
		mid := left + (right-left)/2
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right -= 1
		}

	}
	return numbers[left]
}

func main() {
	test := []int{3, 1, 2, 5}
	fmt.Println(minArray(test))
}
