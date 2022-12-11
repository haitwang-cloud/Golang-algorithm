package main

import "fmt"

//https://leetcode.cn/problems/nZZqjQ/

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, 1
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}
	for left < right {
		mid := left + (right-left)/2
		if canFinish(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
func canFinish(piles []int, h int, k int) bool {
	time := 0
	for _, pile := range piles {
		if pile%k == 0 {
			time += pile / k
		} else {
			time += pile/k + 1
		}
	}
	return time <= h
}

func main() {
	piles, h := []int{30, 11, 23, 4, 20}, 6
	fmt.Println(minEatingSpeed(piles, h))

}
