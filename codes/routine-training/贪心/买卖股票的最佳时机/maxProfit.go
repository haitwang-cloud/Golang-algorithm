package main

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		profit += max(0, prices[i]-prices[i-1])
	}
	return profit
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//买卖股票的最佳时机 II
/*

定义状态
dp[i][0] 表示第 i 天交易完后手里没有股票的最大利润，dp[i][1] 表示第 i 天交易完后手里持有一支股票的最大利润

*/

func maxProfitNew(prices []int) int {
	dp0, dp1 := 0, -prices[0]
	for index := 1; index < len(prices); index++ {
		dp0, dp1 = max(dp0, dp1+prices[index]), max(dp1, dp0-prices[index])
	}
	return dp0
}
