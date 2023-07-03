package main

// https://leetcode.cn/problems/move-zeroes/
func moveZeroes(nums []int) {
	slow := removeElement(nums, 0)
	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
