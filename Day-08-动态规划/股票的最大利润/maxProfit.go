package main

import "fmt"

func maxProfit(prices []int) int {
	maxValue := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxValue += prices[i] - prices[i-1]
		}

	}
	return maxValue
}

func main() {
	test := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(test))
}
