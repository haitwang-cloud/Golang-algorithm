package main

import "fmt"

func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left, right := 1, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return left
}

func main() {
	test := 5
	fmt.Println(mySqrt(test))
}
