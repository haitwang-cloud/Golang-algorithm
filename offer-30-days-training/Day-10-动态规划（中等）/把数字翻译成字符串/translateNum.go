package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {

	src := strconv.Itoa(num)
	a, b := 1, 1
	for index := 2; index <= len(src); index++ {
		subStr := src[index-2 : index]
		var c int
		if subStr <= "25" && subStr >= "10" {
			c = a + b
		} else {
			c = a
		}
		b, a = a, c
	}
	return a
}

func main() {
	test := 1068385902
	fmt.Println(translateNum(test))
}
