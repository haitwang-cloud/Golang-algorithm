package main

import "fmt"

/*

 算法的步骤如下：

    （1）找出待排序的数组中最大和最小的元素
    （2）统计数组中每个值为i的元素出现的次数，存入数组C的第i项
    （3）对所有的计数累加（从C中的第一个元素开始，每一项和前一项相加）
    （4）反向填充目标数组：将每个元素i放在新数组的第C(i)项，每放一个元素就将C(i)减去1

*/
func countingSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	count := make([]int, max-min+1)
	for _, v := range nums {
		count[v-min]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	result := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		index := count[nums[i]-min] - 1
		result[index] = nums[i]
		count[nums[i]-min]--
	}
	return result
}

func main() {
	nums := []int{1, 4, 2, 7, 9, 6, 5, 8, 3, 10}
	fmt.Println(countingSort(nums))
}
