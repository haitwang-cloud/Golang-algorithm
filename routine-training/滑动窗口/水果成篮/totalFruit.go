package main

import (
	"fmt"
)

//https://leetcode.cn/problems/fruit-into-baskets/

func totalFruit(fruits []int) int {

	counter := make(map[int]int)
	start, end, result := 0, 0, 0
	for end < len(fruits) {
		counter[fruits[end]]++
		for start <= end && len(counter) >= 3 {
			counter[fruits[start]]--
			if counter[fruits[start]] == 0 {
				delete(counter, fruits[start])
			}
			start++
		}
		if end-start+1 > result {
			result = end - start + 1
		}
		end++
	}
	return result

}

func main() {
	fruits := []int{1, 2, 1}
	fmt.Println(totalFruit(fruits))
}
