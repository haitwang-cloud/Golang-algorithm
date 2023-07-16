package main

func maxProfit(prices []int, fee int) int {

	/*
		状态转移方程
			dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
			dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i] - fee)
	*/
	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i] - fee
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)

	}
	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
