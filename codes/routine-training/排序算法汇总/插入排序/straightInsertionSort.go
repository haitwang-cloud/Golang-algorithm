package main

import "fmt"

/*
将一个记录插入到已经排好序的有序表中，从而得到一个新的有序表。
通俗理解，我们首先将序列分成两个区间，有序区间和无序区间，
我们每次在无序区间内取一个值，在已排序区间中找到合适的插入位置将其插入，
并保证已排序区间一直有序。下面我们看一下动图吧
最好时间复杂度O(n)
最坏时间复杂度O(n^2)
平均时间复杂度O(n^2)
空间复杂度O(1)
*/

func StraightInsertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		// 从无序区间取出一个值
		if nums[i] < nums[i-1] {
			temp := nums[i]
			j := i - 1
			// 在有序区间找到合适的插入位置
			for ; j >= 0 && nums[j] > temp; j-- {
				nums[j+1] = nums[j]
			}
			nums[j+1] = temp
		}
	}
	return nums

}

func main() {
	nums := []int{1, 4, 2, 7, 9, 6, 5, 8, 3, 10}
	fmt.Println(StraightInsertionSort(nums))
}
