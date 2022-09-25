package main

import (
	"fmt"
)

//链接：https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/solution/jian-zhi-offer-golangbao-mu-xi-lie-ti-ji-ngiz/

func isStraight(nums []int) bool {

	repeat := make(map[int]bool)
	ma, mi := 0, 14
	for _, num := range nums {
		if num == 0 {
			continue
		}
		ma = max(ma, num)
		mi = min(mi, num)
		if _, ok := repeat[num]; ok {
			return false
		}
		repeat[num] = true
	}
	return ma-mi < 5
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y

}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y

}

func main() {
	test := []int{1, 0, 3, 0, 5}
	fmt.Println(isStraight(test))
}
