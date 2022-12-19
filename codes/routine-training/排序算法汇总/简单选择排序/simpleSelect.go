package main

import "fmt"

/*
简单选择排序主要思路就是我们每一趟在 n-i+1 个记录中选取关键字最小的记录作为有序序列的第 i 个记录。
*/

func SimpleSelectSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}

func main() {
	nums := []int{1, 4, 2, 7, 9, 6, 5, 8, 3, 10}
	fmt.Println(SimpleSelectSort(nums))
}
