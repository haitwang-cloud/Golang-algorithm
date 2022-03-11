package main

import (
	"fmt"
)

func findLongestWord(s string, dictionary []string) string {
	maxLen, maxWord := 0, ""
	for _, word := range dictionary {
		if maxLen > len(word) || (maxLen == len(word) && word > maxWord) {
			continue
		}
		if isSubsequence(s, word) {
			maxLen = len(word)
			maxWord = word
		}
	}
	return maxWord
}

func isSubsequence(s1 string, target string) bool {
	i, j := 0, 0
	for i < len(s1) && j < len(target) {
		if s1[i] == target[j] {
			j++
		}
		i++
	}
	return j == len(target)
}

func main() {
	s, dicts := "abce", []string{"abe", "abc"}
	fmt.Println(findLongestWord(s, dicts))
}
