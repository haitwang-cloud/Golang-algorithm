package main

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans, right := 1, intervals[0][1]
	for _, value := range intervals[1:] {
		if value[0] >= right {
			ans++
			right = value[1]
		}
	}
	return n - ans

}

func main() {
	test := [][]int{{3, 4}, {1, 2}, {5, 6}}
	fmt.Println(eraseOverlapIntervals(test))
}
