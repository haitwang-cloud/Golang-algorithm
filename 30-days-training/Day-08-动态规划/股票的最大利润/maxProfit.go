package main

import (
	"fmt"
	"math"
)

// 如果我们真的在买卖股票，我们肯定会想：
// 如果我是在历史最低点买的股票就好了！太好了，
// 在题目中，我们只要用一个变量记录一个历史最低价格 minprice，
// 我们就可以假设自己的股票是在那天买的。那么我们在第 i 天卖出股票能得到的利润就是 prices[i] - minprice

func maxProfit(prices []int) int {
	maxValue, minPrice := 0, math.MaxInt64
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if (prices[i] - minPrice) > maxValue {
			maxValue = prices[i] - minPrice
		}

	}
	return maxValue
}

func main() {
	test := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(test))
}
