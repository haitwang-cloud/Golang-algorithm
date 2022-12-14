package main

import "fmt"

//https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/solution/mian-shi-ti-53-ii-0n-1zhong-que-shi-de-shu-zi-er-f/
/*

左子数组：nums[index]=index
右子数组: nums[index]!=index
所以此题目是查找右字数组的边界
*/
func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return left
}

func main() {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println(missingNumber(test))
}
