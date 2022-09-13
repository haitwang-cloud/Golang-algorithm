package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	maxSum := math.MinInt32
	if len(nums) == 0 {
		return maxSum
	}
	curSum := 0
	for index := 0; index < len(nums); index++ {
		if curSum > 0 {
			curSum += nums[index]
		} else {
			curSum = nums[index]
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}

func main() {
	test := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(test))
}
