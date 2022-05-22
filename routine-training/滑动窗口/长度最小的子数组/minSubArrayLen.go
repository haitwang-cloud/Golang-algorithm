package main

import (
	"fmt"
	"math"
)

//https://leetcode.cn/problems/minimum-size-subarray-sum/

func minSubArrayLen(target int, nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	start, end, sum, result := 0, 0, 0, math.MaxInt32
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			subLength := end - start + 1
			if result > subLength {
				result = subLength
			}
			sum -= nums[start]
			start++
		}
		end++
	}
	if result == math.MaxInt32 {
		return 0
	}
	return result

}

func main() {
	nums, test := []int{2, 3, 1, 2, 4, 3}, 7
	fmt.Println(minSubArrayLen(test, nums))
}
