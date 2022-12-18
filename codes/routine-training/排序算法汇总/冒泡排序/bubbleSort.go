package main

import "fmt"

/*
	冒泡排序不断进行交换，通过交换完成最终的排序
	最好时间复杂度O(n)
	最坏时间复杂度O(n^2)
	平均时间复杂度O(n^2)
	空间复杂度O(1)
*/
func BubbleSortPro(nums []int) []int {
	length := len(nums)
	var flag bool
	for i := 0; i < length && flag; i++ {
		for j := 0; j < length-i-1; j++ {
			//如果没发生交换，则依旧为false，下次就会跳出循环
			flag = false
			if nums[j] > nums[j+1] {
				//发生交换，则变为true,下次继续判断
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
	}
	return nums
}

func BubbleSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	nums := []int{1, 4, 2, 7, 9, 6, 5, 8, 3, 10}
	fmt.Println(BubbleSort(nums))
}
