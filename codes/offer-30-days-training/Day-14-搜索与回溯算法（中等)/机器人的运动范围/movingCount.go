package main

import "fmt"

/*
链接：https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/solution/can-kao-krahetsda-shen-de-jie-fa-de-goyu-fgja/
*/

var n1, m1, k1 int
var visited [][]bool

func movingCount(m int, n int, k int) int {
	m1 = m
	n1 = n
	k1 = k
	visited = make([][]bool, m)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, n)
	}

	return dfs(0, 0, 0, 0)
}

func dfs(i int, j int, si int, sj int) int {
	if i >= m1 || j >= n1 || si+sj > k1 || visited[i][j] {
		return 0
	}
	visited[i][j] = true

	sj1 := sj + 1
	si1 := si + 1
	if (j+1)%10 == 0 {
		sj1 = sj - 8
	}
	if (i+1)%10 == 0 {
		si1 = si - 8
	}

	return 1 + dfs(i, j+1, si, sj1) + dfs(i+1, j, si1, sj)
}

func main() {
	m, n, k := 2, 3, 1
	fmt.Println(movingCount(m, n, k))
}
