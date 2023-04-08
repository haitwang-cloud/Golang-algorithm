package main

func twoSum(numbers []int, target int) []int {

	// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/
	if len(numbers) < 2 {
		return []int{-1, -1}
	}
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
