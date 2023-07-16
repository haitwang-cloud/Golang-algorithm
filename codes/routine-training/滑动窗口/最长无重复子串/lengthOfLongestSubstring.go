package main

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	res := 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		// 判断左侧窗口是否要收缩
		for window[c] > 1 {
			d := s[left]
			left++
			// 进行窗口内数据的一系列更新
			window[d]--
		}
		// 在这里更新答案
		res = max(res, right-left)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
