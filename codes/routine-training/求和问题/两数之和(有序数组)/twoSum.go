package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{-1, -1}

}

func main() {

	numbers, target := []int{3, 2, 4}, 6
	fmt.Println(twoSum(numbers, target))

}
