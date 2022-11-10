package main

import "fmt"

//链接：https: //leetcode.cn/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof/solution/go-shou-xi-bian-cheng-mei-sha-nan-du-by-ba-xiang/

func printNumbers(n int) []int {

	var res []int
	var max int
	for n > 0 {
		max = max*10 + 9
		n--
	}
	for i := 1; i <= max; i++ {
		res = append(res, i)
	}
	return res

}
func main() {
	test := 3
	fmt.Println(printNumbers(test))

}
