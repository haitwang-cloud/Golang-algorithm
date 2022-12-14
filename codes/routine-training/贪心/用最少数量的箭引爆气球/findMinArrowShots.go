package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	ans := 1
	maxRight := points[0][1]
	for _, value := range points {
		if value[0] > maxRight {
			maxRight = value[1]
			ans++
		}
	}
	return ans

}

func main() {
	test := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(test))
}
