package fourSum

import "sort"

func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	right := len(nums)
	for i := 0; i < right; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		threeSumRes := threeSumTarget(nums, i+1, target-nums[i])
		for _, pair := range threeSumRes {
			res = append(res, []int{nums[i], pair[0], pair[1], pair[2]})
		}
	}
	return res
}

func threeSumTarget(nums []int, start, target int) [][]int {
	res := make([][]int, 0)
	right := len(nums)
	for i := start; i < right; i++ {
		if i > start && nums[i] == nums[i-1] {
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
