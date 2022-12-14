package main

import (
	"fmt"
	"math"
	"sort"
)

// https://leetcode.cn/problems/heaters/

func findRadius(houses []int, heaters []int) int {
	ans := 0
	sort.Ints(heaters)
	for _, house := range houses {
		j := sort.SearchInts(heaters, house+1)
		minDis := math.MaxInt32
		if j < len(heaters) {
			minDis = heaters[j] - house
		}
		i := j - 1
		if i >= 0 && house-heaters[i] < minDis {
			minDis = house - heaters[i]
		}
		if minDis > ans {
			ans = minDis
		}
	}
	return ans
}

func main() {

	nums, test := []int{1, 2, 3}, []int{2}
	fmt.Println(findRadius(nums, test))

}
