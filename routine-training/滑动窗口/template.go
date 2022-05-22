package main

func slideWindow(nums []int) int {
	counter := make(map[int]int)
	start, end, result := 0, 0, 0
	for end < len(nums) {
		counter[nums[end]]++
		for start <= end && check() {

			start++
		}
		end++
	}
	return result

}
