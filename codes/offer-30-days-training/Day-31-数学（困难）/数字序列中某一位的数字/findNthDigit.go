package main

import (
	"fmt"
	"strconv"
)

//https://leetcode.cn/problems/shu-zi-xu-lie-zhong-mou-yi-wei-de-shu-zi-lcof/solution/mian-shi-ti-44-shu-zi-xu-lie-zhong-mou-yi-wei-de-6/

func findNthDigit(n int) int {

	digit, digitNum, count := 1, 1, 9
	for n > count {
		n -= count
		digit++
		digitNum *= 10
		count = 9 * digit * digitNum
	}
	num := digitNum + (n-1)/digit

	index := (n - 1) % digit

	numStr := strconv.Itoa(num)
	return int(numStr[index] - '0')

}

func main() {
	test := 3
	fmt.Println(findNthDigit(test))

}
