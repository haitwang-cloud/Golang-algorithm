package main

import (
	"fmt"
)

/*
比较特别的情况就是 nums [i] < nums [i - 2]，
修改 nums [i - 1] = nums [i] 不能使数组成为非递减数组，
只能修改 nums[i+1] = nums[i]
*/

func checkPossibility(nums []int) bool {
	cnt := 0
	for index := 0; index < len(nums)-1; index++ {
		if nums[index] > nums[index+1] {
			cnt++
			if cnt > 1 {
				return false
			}
			if index > 0 && nums[index+1] < nums[index-1] {
				nums[index+1] = nums[index]
			}
		}

	}
	return true

}

func main() {
	people := []int{4, 2, 3}
	fmt.Println(checkPossibility(people))
}
