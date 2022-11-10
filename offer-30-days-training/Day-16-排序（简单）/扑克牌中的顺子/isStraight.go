package main

import (
	"fmt"
)

//链接：https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/solution/jian-zhi-offer-golangbao-mu-xi-lie-ti-ji-ngiz/
/*

统计王的数目，joker值就是第一个不是王的元素的下标
如果两个非零元素相等，那么返回false
最后判断，最大值（排序后）和最小值的差值，只要<5那么就可以组成顺组，比如：0 0 1 3 5，0 0 3 5 7，只要小于五，五个位置可以全部补齐
*/
func isStraight(nums []int) bool {

	repeatDict := make(map[int]bool)
	maxValue, minValue := 0, 14
	for _, value := range nums {
		if value == 0 {
			continue
		}
		minValue = min(minValue, value)
		maxValue = max(maxValue, value)
		if _, ok := repeatDict[value]; ok {
			return false
		}
		repeatDict[value] = true
	}
	return maxValue-minValue < 5
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
