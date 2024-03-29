package main

import "fmt"

func dicesProbability(n int) []float64 {

	dp := make([][]float64, n)
	for i := range dp {
		dp[i] = make([]float64, (i+1)*6-i)
	}
	for i := range dp[0] {
		dp[0][i] = float64(1) / float64(6)
	}
	for i := 1; i < len(dp); i++ {
		for j := range dp[i-1] {
			for k := range dp[0] {
				dp[i][j+k] += float64(dp[i-1][j]) * float64(dp[0][k])
			}
		}
	}
	return dp[n-1]

	//链接：https://leetcode.cn/problems/nge-tou-zi-de-dian-shu-lcof/solution/golanger-wei-dong-tai-gui-hua-jian-ji-yi-dv8q/。
	
}

func main() {
	test := 2
	fmt.Println(dicesProbability(test))
}
