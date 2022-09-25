package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		bit := 0 //记录该位上的和
		for _, num := range nums {
			bit += (num >> i) & 1
		}
		res += (bit % 3) << i
	}
	return res
}

func main() {

	test := []int{4, 4, 4, 6}
	fmt.Println(singleNumber(test))
}
