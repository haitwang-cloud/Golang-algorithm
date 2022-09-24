package main

import "fmt"

//https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/solution/mian-shi-ti-15-er-jin-zhi-zhong-1de-ge-shu-wei-yun/

func hammingWeight(num uint32) int {
	count := 0
	for ; num > 0; num &= num - 1 {
		count++
	}
	return count
}

func main() {
	test := uint32(00000000000000000000000000001011)
	fmt.Println(hammingWeight(test))
}
