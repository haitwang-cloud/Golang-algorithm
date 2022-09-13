package main

import (
	"fmt"
	"math"
)

// 在题目中，我们只要用一个变量记录一个历史最低价格 minprice，
// 我们就可以假设自己的股票是在那天买的。那么我们在第 i 天卖出股票能得到的利润就是 prices[i] - minprice
// profit = max(profit, price - minPrice)

func maxProfit(prices []int) int {
	minPrice, profit := math.MaxInt64, 0
	for index := 0; index < len(prices); index++ {
		if minPrice > prices[index] {
			minPrice = prices[index]
		}
		if prices[index]-minPrice > profit {
			profit = prices[index] - minPrice
		}
	}

	return profit

}

func main() {
	test := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(test))
}
