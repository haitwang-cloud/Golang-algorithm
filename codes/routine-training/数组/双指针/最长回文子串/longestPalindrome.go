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
		s1, s2 := palindrome(s, i, i), palindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}
