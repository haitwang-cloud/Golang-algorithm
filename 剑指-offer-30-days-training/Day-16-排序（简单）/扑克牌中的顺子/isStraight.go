package main

import (
	"fmt"
	"sort"
)

//链接：https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/solution/jian-zhi-offer-golangbao-mu-xi-lie-ti-ji-ngiz/
func isStraight(nums []int) bool {
	sort.Ints(nums)
	joker := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			joker++ //记录joker数量
		} else {
			if nums[i] == nums[i+1] { //除了0之外，有相同的元素直接返回false
				return false
			}
		}
	}
	return nums[4]-nums[joker] < 5
}

func isStraightNew(nums []int) bool {

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

//def isStraight(self, nums: List[int]) -> bool:
//repeat = set()
//ma, mi = 0, 14
//for num in nums:
//if num == 0: continue # 跳过大小王
//ma = max(ma, num) # 最大牌
//mi = min(mi, num) # 最小牌
//if num in repeat: return False # 若有重复，提前返回 false
//repeat.add(num) # 添加牌至 Set
//return ma - mi < 5 # 最大牌 - 最小牌 < 5 则可构成顺子

func main() {
	test := []int{1, 2, 3, 5, 5}
	fmt.Println(isStraightNew(test))
}
