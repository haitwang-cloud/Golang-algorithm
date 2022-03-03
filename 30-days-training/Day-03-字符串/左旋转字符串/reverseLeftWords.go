package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	result_left, result_right := s[:n], s[n:]
	return result_right + result_left

}
func reverseLeftWordsNew(s string, n int) string {
	str_result := []byte(s)
	// 1. 反转前n个字符
	// 2. 反转第n到end字符
	// 3. 反转整个字符
	reverse(str_result, 0, n-1)
	reverse(str_result, n, len(str_result)-1)
	reverse(str_result, 0, len(str_result)-1)
	return string(str_result)

}
func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}

}

func main() {
	test := "abcdefg"
	fmt.Println(reverseLeftWords(test, 2))
	fmt.Println(reverseLeftWordsNew(test, 2))
}
