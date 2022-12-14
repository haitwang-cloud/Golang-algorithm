package main

import (
	"fmt"
	"sort"
)

/*

https://leetcode-cn.com/problems/queue-reconstruction-by-height/solution/gen-ju-shen-gao-zhong-jian-dui-lie-by-leetcode-sol/
*/

func reconstructQueue(people [][]int) [][]int {
	ans := make([][]int, len(people))

	if len(people) == 0 {
		return ans
	}

	sort.Slice(people, func(i, j int) bool {

		return people[i][0] < people[j][0] || (people[i][0] == people[j][0] && people[i][1] > people[j][1])
	})
	fmt.Println(people)
	for _, person := range people {
		spaces := person[1] + 1
		for index := range ans {
			if ans[index] == nil {
				spaces--
				if spaces == 0 {
					ans[index] = person
					break

				}

			}
		}
	}
	return ans

}

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(people))
}
