package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	result_left, result_right := s[:n], s[n:]
	return result_right + result_left

}

func main() {
	test := "abcdefg"
	fmt.Println(reverseLeftWords(test, 2))
}
