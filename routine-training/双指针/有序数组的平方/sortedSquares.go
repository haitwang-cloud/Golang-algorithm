package main

import "fmt"

func sortedSquares(nums []int) []int {

	ans := make([]int, len(nums))
	left, right, index := 0, len(nums)-1, len(nums)-1
	for left <= right {
		if nums[left]*nums[left] >= nums[right]*nums[right] {
			ans[index] = nums[left] * nums[left]
			left++
		} else {
			ans[index] = nums[right] * nums[right]
			right--
		}
		index--
	}

	return ans

}

func main() {
	test := []int{-7, -3, 2, 3, 11}
	//sortColors(test)
	fmt.Println(sortedSquares(test))
}
