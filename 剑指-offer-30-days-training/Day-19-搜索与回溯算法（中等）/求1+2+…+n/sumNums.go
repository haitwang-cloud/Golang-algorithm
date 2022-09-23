package main

import "fmt"

func sumNums(n int) int {

	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}

func main() {
	test := 5
	fmt.Println(sumNums(test))

}
