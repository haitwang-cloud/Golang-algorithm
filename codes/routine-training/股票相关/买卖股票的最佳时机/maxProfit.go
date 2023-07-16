package main

import "math"

func maxProfit(prices []int) int {

	/* 状态转移方程
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
	    dp[i][1] = max(dp[i-1][1], -prices[i])
	*/

	n := len(prices)
	/*
		dp := make([][2]int, n)
		for i := 0; i < n; i++ {
			if i-1 == -1 {
				// base case
				dp[i][0] = 0
				dp[i][1] = -prices[i]
				continue
			}
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], -prices[i])
		}
		return dp[n-1][0]
	*/
	//空间优化版本
	dp_i_0, dp_i_1 := 0, math.MinInt32
	for i := 0; i < n; i++ {
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, -prices[i])
	}
	return dp_i_0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
