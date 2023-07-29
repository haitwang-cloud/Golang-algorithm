package threesum

import "sort"

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	return threeSumTarget(nums, 0)
}

func threeSumTarget(nums []int, target int) [][]int {
	res := make([][]int, 0)
	right := len(nums)
	for i := 0; i < right; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		twoRes := twoSum(nums, i+1, target-nums[i])
		for _, pair := range twoRes {
			res = append(res, []int{nums[i], pair[0], pair[1]})
		}
	}
	return res
}

func twoSum(numbers []int, start, target int) [][]int {
	res := make([][]int, 0)
	if len(numbers) < 2 {
		return res
	}
	left, right := start, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			res = append(res, []int{numbers[left], numbers[right]})
			for left < right && numbers[left] == numbers[left+1] {
				left++
			}
			for left < right && numbers[right] == numbers[right-1] {
				right--
			}
			left++
			right--
		}
	}
	return res
}
