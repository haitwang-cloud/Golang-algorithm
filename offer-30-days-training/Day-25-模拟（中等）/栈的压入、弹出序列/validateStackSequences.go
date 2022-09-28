package main

import "fmt"

//链接：https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/solution/mian-shi-ti-31-zhan-de-ya-ru-dan-chu-xu-lie-mo-n-2/
func validateStackSequences(pushed []int, popped []int) bool {
	var stack []int
	index := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[index] {
			stack = stack[:len(stack)-1]
			index++
		}
	}
	return len(stack) == 0
}

func main() {
	pushed, poped := []int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(pushed, poped))
}
