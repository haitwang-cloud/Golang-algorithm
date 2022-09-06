package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	left, right := s[:n], s[n:]
	return right + left

}
func reverseLeftWordsNew(s string, n int) string {
	result := []byte(s)
	// 1. 反转前n个字符
	// 2. 反转第n到end字符
	// 3. 反转整个字符
	reverse(result, 0, n-1)
	reverse(result, n, len(result)-1)
	reverse(result, 0, len(result)-1)
	return string(result)

}
func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left += 1
		right -= 1
	}

}

func main() {
	test := "abcdefg"
	fmt.Println(reverseLeftWords(test, 2))
	fmt.Println(reverseLeftWordsNew(test, 2))
}
