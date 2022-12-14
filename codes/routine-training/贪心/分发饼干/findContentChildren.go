package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	if len(g) == 0 || len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}

func main() {
	g, s := []int{1, 2, 3}, []int{1, 1, 2}
	fmt.Println(findContentChildren(g, s))
}
