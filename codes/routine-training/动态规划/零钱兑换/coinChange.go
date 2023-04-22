package main

// dp 数组的定义：当目标金额为 i 时，至少需要 dp[i] 枚硬币凑出。
func coinChange(coins []int, amount int) int {
	// 初始化dp数组为amount+1
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	// 初始化dp[0]为0
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		// 遍历每个硬币的面值
		for _, coin := range coins {
			if coin <= i {
				// 计算dp[i]的值
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}

// 求两个数的最小值
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
