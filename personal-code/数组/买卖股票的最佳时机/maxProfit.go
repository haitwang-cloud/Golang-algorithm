package main

import "fmt"

func maxProfit(prices []int) (profit int) {
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
func main() {
	test := []int{2, 3, 1, 5, 6}
	fmt.Println(maxProfit(test))
}
