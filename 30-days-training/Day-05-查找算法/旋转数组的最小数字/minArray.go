package main

import "fmt"

func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	left, right := 0, len(numbers)-1
	for left < right {
		mid := left + (right-left)/2
		if numbers[mid] > numbers[right] {
			left = mid + 1
		} else if numbers[mid] < numbers[right] {
			right = mid
		} else {
			right -= 1
		}

	}
	return numbers[left]
}

func main() {
	test := []int{3, 1, 2, 5}
	fmt.Println(minArray(test))
}
