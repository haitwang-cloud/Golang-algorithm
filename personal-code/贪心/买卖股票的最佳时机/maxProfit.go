package main

import "fmt"

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

func maxProfitNew(prices []int) int {
	n := len(prices)
	dp0, dp1 := 0, -prices[0]
	for i := 1; i < n; i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])

	}
	return dp0
}
func main() {
	test := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(test))
	fmt.Println(maxProfitNew(test))
}
