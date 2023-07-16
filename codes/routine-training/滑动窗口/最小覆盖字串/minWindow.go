package main

import "math"

//https://leetcode.cn/problems/minimum-window-substring/
func minWindow(s string, t string) string {
	//初始化 need和slideWindow为 map
	need := make(map[byte]int)   //need为锁需要的字符数
	window := make(map[byte]int) // window是当前滑动串口内的字符数
	for val := range t {
		need[t[val]]++
	}
	start, length := 0, math.MaxInt32 // 最小覆盖子串的起始索引及长度
	//使用 left 和 right 变量初始化窗口的两端
	// 其中 valid 变量表示窗口中满足 need 条件的字符个数，
	//如果 valid 和 need.size 的大小相同，则说明窗口已满足条件，已经完全覆盖了串 T
	left, right, valid := 0, 0, 0
	for right < len(s) {
		c := s[right] // c为当前要加入的字符
		right++
		if _, ok := need[c]; ok { //如果当前字符出现在need中，证明被t需要
			window[c]++               //加入到当前的窗口中
			if window[c] == need[c] { //如果窗口中该字符串的数量等于need中的数量
				valid++ // valid 计数器加 1
			}
		}

		//接下来处理滑动窗口中的字符已经完全覆盖字串 t 中的字符
		for valid == len(need) {
			if length > right-left { // 如果有长度更短的字串出现
				start = left
				length = right - left
			}
			d := s[left] //d为当前要移除的字符
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] { //当前d已经满足了他在字串 t 中的需求
					valid-- //计数器 valid 减一
				}
				window[d]-- //移出窗口
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
