package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}
func twoSumNew(nums []int, target int) []int {
	arrayMap := make(map[int]int)
	for index, value := range nums {
		if item, ok := arrayMap[target-value]; ok {
			return []int{item + 1, index + 1}
		}
		arrayMap[value] = index
	}

	return []int{-1, -1}
}

func main() {
	test, k := []int{2, 7, 11, 15}, 9
	fmt.Println(twoSum(test, k))
	fmt.Println(twoSumNew(test, k))
}
