package main

func rob(nums []int) int {

	memo := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = -1
	}

	return dp(nums, memo, 0)
}

func dp(nums []int, memo []int, start int) int {
	if start >= len(nums) {
		return 0
	}
	res := 0
	if memo[start] != -1 {
		return memo[start]
	}
	res = max(
		dp(nums, memo, start+1),             //当前不抢，直接去下家
		dp(nums, memo, start+2)+nums[start]) //当前抢，直接去下下家
	memo[start] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
