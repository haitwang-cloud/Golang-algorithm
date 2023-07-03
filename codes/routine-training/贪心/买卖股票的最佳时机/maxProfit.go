package main

import "math"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {

	minPrice, maxValue := math.MaxInt64, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if maxValue < prices[i]-minPrice {
			maxValue = prices[i] - minPrice
		}
	}
	return maxValue
}
