package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	preSum := nums[0]
	maxSum := preSum
	for index := 1; index < len(nums); index++ {
		if preSum > 0 {
			preSum += nums[index]

		} else {
			preSum = nums[index]
		}
		if preSum > maxSum {
			maxSum = preSum
		}

	}

	return maxSum

}

func main() {
	people := []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(people))
}
