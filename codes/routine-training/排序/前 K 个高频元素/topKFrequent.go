package main

import (
	"fmt"
	"math/rand"
	"time"
)

// https://leetcode.cn/problems/top-k-frequent-elements/solution/qian-k-ge-gao-pin-yuan-su-by-leetcode-solution/

func topKFrequent(nums []int, k int) []int {
	//统计每个数字出现的次数
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	//将数字和出现的次数放入二维数组中
	values := make([][]int, 0)
	for key, value := range count {
		values = append(values, []int{key, value})
	}
	//将每个数字和出现的次数放入二维数组中
	ret := make([]int, k)
	//对二维数组进行快速排序
	quickSort(values, 0, len(values)-1, ret, 0, k)
	return ret
}

func quickSort(values [][]int, left, right int, ret []int, retIndex, k int) {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int()%(right-left+1) + left
	values[picked], values[left] = values[left], values[picked]
	// 以 values[left][1] 为基准，将数组分为两部分
	pivot, pivotIndex := values[left][1], left
	// 从左到右遍历数组
	for i := left + 1; i <= right; i++ {
		// 如果当前元素的出现次数大于基准值，就将其交换到左侧
		if values[i][1] >= pivot {
			values[pivotIndex+1], values[i] = values[i], values[pivotIndex+1]
			pivotIndex++
		}
	}
	// 将基准值放到中间
	values[left], values[pivotIndex] = values[pivotIndex], values[left]
	// 如果基准值的下标刚好等于 k-1，说明基准值就是第 k 大的元素
	if pivotIndex == k-1 {
		for i := 0; i < k; i++ {
			ret[i] = values[i][0]
		}
		return
	}
	// 如果基准值的下标大于 k-1，说明第 k 大的元素在基准值的左侧
	if pivotIndex > k-1 {
		quickSort(values, left, pivotIndex-1, ret, retIndex, k)
	} else {
		// 如果基准值的下标小于 k-1，说明第 k 大的元素在基准值的右侧
		for i := retIndex; i <= pivotIndex; i++ {
			ret[i] = values[i][0]
		}
		quickSort(values, pivotIndex+1, right, ret, pivotIndex+1, k)
	}
}

func topKFrequentNew(nums []int, k int) (ans []int) {
	// 确定每个数的频率
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// 将不同频率的元素放入不同桶
	buckts := make([][]int, len(nums)+1) // 创建桶
	for k, v := range freq {
		buckts[v] = append(buckts[v], k)
	}

	// 从最大桶开始逐步取出高频元素
	for i := len(buckts) - 1; k > 0; i-- {
		for _, v := range buckts[i] {
			ans = append(ans, v)
			k--
		}
	}
	return
}

// 链接：https://leetcode.cn/problems/top-k-frequent-elements/solution/bi-guan-fang-ti-jie-geng-kuai-de-fang-fa-lgsc/

func main() {
	nums1, k := []int{1, 1, 1, 2, 2, 3}, 2
	fmt.Println(topKFrequent(nums1, k))
	fmt.Println(topKFrequentNew(nums1, k))
}
