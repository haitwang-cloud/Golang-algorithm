package main

func longestPalindrome(s string) string {
	/*
			for 0 <= i < len(s):
		    找到以 s[i] 为中心的回文串
		    找到以 s[i] 和 s[i+1] 为中心的回文串
		    更新答案
	*/
	res := ""
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}

	return res

}

func palindrome(s string, left, right int) string {

	for left >= 0 && right < len(s) && s[left] == s[right] {
		// 向两边展开
		left--
		right++
	}
	return s[left+1 : right]
}
