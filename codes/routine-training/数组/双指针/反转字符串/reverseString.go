package main

// https://leetcode.cn/problems/reverse-string/submissions/
func reverseString(s []byte) {
	if len(s) < 2 {
		return
	}
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
