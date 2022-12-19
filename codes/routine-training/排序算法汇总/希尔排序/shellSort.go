package main

import "fmt"

// https://www.runoob.com/w3cnote/shell-sort.html

/*

希尔排序，也称递减增量排序算法，是插入排序的一种更高效的改进版本。但希尔排序是非稳定排序算法。

希尔排序是基于插入排序的以下两点性质而提出改进方法的：

- 插入排序在对几乎已经排好序的数据操作时，效率高，即可以达到线性排序的效率；
- 但插入排序一般来说是低效的，因为插入排序每次只能将数据移动一位；

希尔排序的基本思想是：先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，
待整个序列中的记录"基本有序"时，再对全体记录进行依次直接插入排序。

*/

func ShellSort(nums []int) []int {
	// 增量gap，并逐步缩小增量
	for gap := len(nums) / 2; gap > 0; gap /= 2 {
		// 从第gap个元素，逐个对其所在组进行直接插入排序操作
		for i := gap; i < len(nums); i++ {
			j := i
			for j-gap >= 0 && nums[j] < nums[j-gap] {
				nums[j], nums[j-gap] = nums[j-gap], nums[j]
				j -= gap
			}
		}
	}
	return nums

}

func main() {
	nums := []int{1, 4, 2, 7, 9, 6, 5, 8, 3, 10}
	fmt.Println(ShellSort(nums))
}
