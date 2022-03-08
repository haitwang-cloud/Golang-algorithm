package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	s1 := []byte(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < len(s) && !strings.Contains("aeiouAEIOU", string(s1[left])) {
			left++
		}
		for right > 0 && !strings.Contains("aeiouAEIOU", string(s1[right])) {
			right--
		}
		if left < right {
			s1[left], s1[right] = s1[right], s1[left]
			left++
			right--
		}
	}
	return string(s1)

}

func main() {
	test := "leetcode"
	fmt.Println(reverseVowels(test))
}
