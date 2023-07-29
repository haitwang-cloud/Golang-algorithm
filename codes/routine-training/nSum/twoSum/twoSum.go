package twoSum

func twoSum(numbers []int, target int) [][]int {
	if len(numbers) < 2 {
		return [][]int{}
	}
	res := [][]int{}
	left, right := 0, len(numbers)-1
	for left < right {
		l, r := numbers[left], numbers[right]
		sum := l + r
		if sum > target {
			for left < right && l == numbers[left] {
				left++
			}
		} else if sum < target {
			for left < right && r == numbers[right] {
				right--
			}
		} else {
			res = append(res, []int{left, right})
			for left < right && l == numbers[left] {
				left++
			}
			for left < right && r == numbers[right] {
				right--
			}

		}
	}
	return res
}
