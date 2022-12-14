package main

import "fmt"

//https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/solution/jian-zhi-offer-40-zui-xiao-de-k-ge-shu-j-9yze/
//https://leetcode.cn/problems/zui-xiao-de-kge-shu-lcof/solution/dui-pai-gui-bing-kuai-pai-shi-xian-zui-x-34vb/

func getLeastNumbers(arr []int, k int) []int {
	//quicksort(arr, 0, len(arr)-1)
	//return arr[:k]

	return quicksearch(arr, 0, len(arr)-1, k)

}

func partition(nums []int, i, j int) int {
	l, m, r := i, i, j
	for l < r {
		for l < r && nums[r] >= nums[m] {
			r--
		}
		for l < r && nums[l] <= nums[m] {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[m], nums[l] = nums[l], nums[m]

	return l
}

func quicksort(nums []int, i, j int) {
	if i >= j {
		return
	}

	m := partition(nums, i, j)
	quicksort(nums, i, m-1)
	quicksort(nums, m+1, j)
}

func quicksearch(nums []int, i, j, k int) []int {

	t := partition(nums, i, j)

	if t == k-1 {
		return nums[:k]
	}

	if t < k-1 {
		return quicksearch(nums, t+1, j, k)
	}

	return quicksearch(nums, i, t-1, k)

}

func main() {
	test, k := []int{1, 0, 3, 0, 5}, 2
	fmt.Println(getLeastNumbers(test, k))
}
