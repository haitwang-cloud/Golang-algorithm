package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	hashMap := make(map[int]int, len(numbers))

	for i, num := range numbers {
		if index, ok := hashMap[target-num]; ok {
			return []int{index, i}
		}
		hashMap[num] = i
	}
	return []int{-1, -1}
}

func main() {

	numbers, target := []int{3, 2, 4}, 6
	fmt.Println(twoSum(numbers, target))

}
