package main

import "fmt"

//https://leetcode.cn/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/solution/mian-shi-ti-65-bu-yong-jia-jian-cheng-chu-zuo-ji-7/

func add(a int, b int) int {

	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a

}

func main() {
	a, b := 1, 2
	fmt.Println(add(a, b))
}
