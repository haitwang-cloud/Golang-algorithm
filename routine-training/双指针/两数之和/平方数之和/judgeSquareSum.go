package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sums := left*left + right*right
		if sums == c {
			return true
		} else if sums > c {
			right--
		} else {
			left++
		}
	}
	return false
}

func main() {
	test := 5
	fmt.Println(judgeSquareSum(test))
}
