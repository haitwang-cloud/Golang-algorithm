package main

func removeDuplicates(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			// 维护 nums[0..slow] 无重复
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
