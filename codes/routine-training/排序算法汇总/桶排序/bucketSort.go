package main

import "fmt"

var DEFALUT_BUCKET_SIZE = 5

// https://www.runoob.com/w3cnote/bucket-sort.html
func bucketSort(nums []int) []int {

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
	bucketCount := (max-min)/DEFALUT_BUCKET_SIZE + 1
	buckets := make([][]int, bucketCount)
	for i := 0; i < bucketCount; i++ {
		buckets[i] = make([]int, 0)
	}
	//利用映射函数将数据分配到各个桶中
	for _, v := range nums {
		buckets[(v-min)/DEFALUT_BUCKET_SIZE] = append(buckets[(v-min)/DEFALUT_BUCKET_SIZE], v)
	}
	result := make([]int, 0)
	for _, v := range buckets {
		if len(v) <= 1 {
			result = append(result, v...)
		} else {
			// 对每个桶进行排序，这里使用了插入排序
			insertSort(v)
			result = append(result, v...)
		}
	}
	return result
}

func insertSort(nums []int) {
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
}

func main() {
	nums := []int{1, 4, 2, 7, 9, 6, 5, 8, 3, 10}
	fmt.Println(bucketSort(nums))
}
