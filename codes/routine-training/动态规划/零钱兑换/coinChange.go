package main

/*
假设你有面值为 1, 2, 5 的硬币，你想求 amount = 11 时的最少硬币数（原问题）
，如果你知道凑出 amount = 10, 9, 6 的最少硬币数（子问题），
你只需要把子问题的答案加一（再选一枚面值为 1, 2, 5 的硬币），
求个最小值，就是原问题的答案。因为硬币的数量是没有限制的，所以子问题之间没有相互制，是互相独立的。

*/
func coinChange(coins []int, amount int) int {
	// 初始化dp数组为amount+1
	// dp 数组的定义：当目标金额为 i 时，至少需要 dp[i] 枚硬币凑出。
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	// 外层 for 循环在遍历所有状态的所有取值
	for i := 0; i < len(dp); i++ {
		// 内层 for 循环在求所有选择的最小值
		for _, coin := range coins {
			// 子问题无解，跳过
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

// 求两个数的最小值
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
