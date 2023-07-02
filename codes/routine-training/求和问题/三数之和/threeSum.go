package main

import (
	"sort"
)

// https://leetcode.cn/problems/3sum/solution/hua-jie-suan-fa-15-san-shu-zhi-he-by-guanpengchn/

/*
	排序+双指针
	通过排序，可以将问题转化为两数之和的问题，从而可以使用双指针的方法解决。
	首先对数组进行排序，排序后固定一个数 nums[i]，再使用左右指针指向 nums[i]后面的两端，
	数字分别为 nums[L] 和 nums[R]，
	计算三个数的和 sum 判断是否满足为 0，满足则添加进结果集

	如果 nums[i]大于 0，则三数之和必然无法等于 0，结束循环
	如果 nums[i] == nums[i-1]，则说明该数字重复，会导致结果重复，所以应该跳过
	如果 sum == 0，nums[L] == nums[L+1] 则会导致结果重复，应该跳过，L++
	如果 sum == 0，nums[R] == nums[R-1] 则会导致结果重复，应该跳过，R--

*/

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)
	for first := 0; first < len(nums)-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		left, right := first+1, len(nums)-1
		for left < right {
			sum := nums[first] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[first], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return ans
}
