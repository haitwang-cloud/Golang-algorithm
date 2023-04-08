package main

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			/*
				注意这里和有序数组去重的解法有一个细节差异，我们这里是先给 nums[slow] 赋值然后再给 slow++，
				这样可以保证 nums[0..slow-1] 是不包含值为 val 的元素的，最后的结果数组长度就是 slow。
			*/
			slow++
		}
		fast++
	}
	return slow

}
