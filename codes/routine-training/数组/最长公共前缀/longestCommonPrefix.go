package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, value := range strs {
		for strings.Index(value, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func longestCommonPrefixNew(strs []string) string {
	//分治

	if len(strs) == 0 {
		return ""
	}
	var lcp func(int, int) string
	lcp = func(start int, end int) string {
		if start == end {
			return strs[start]
		}
		mid := (start + end) / 2
		lcpLeft, lcpRight := lcp(start, mid), lcp(mid+1, end)
		minLength := min(len(lcpLeft), len(lcpRight))
		for index := 0; index < minLength; index++ {
			if lcpLeft[index] != lcpRight[index] {
				return lcpLeft[:index]
			}
		}
		return lcpLeft[:minLength]

	}

	return lcp(0, len(strs)-1)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	test := []string{"flower", "flow", "flight", "flip"}
	fmt.Println(longestCommonPrefix(test))
	fmt.Println(longestCommonPrefixNew(test))
}
