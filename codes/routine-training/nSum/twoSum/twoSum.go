package twoSum

func twoSum(numbers []int, target int) [][]int {
	res := make([][]int, 0)
	if len(numbers) < 2 {
		return res
	}
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum > target {
			right--
		} else if sum < target {
			left++
		} else {
			res = append(res, []int{left, right})
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
