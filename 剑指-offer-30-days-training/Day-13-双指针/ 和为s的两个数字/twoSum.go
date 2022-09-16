package main

import "fmt"

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == target {
			return []int{nums[left], nums[right]}
		} else if nums[left]+nums[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{-1, -1}

}

func main() {
	test, k := []int{2, 7, 11, 15}, 9
	fmt.Println(twoSum(test, k))
}
