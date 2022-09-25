package main

import "fmt"

/*

记数组nums 的长度为n。先从nums 左侧开始遍历，
如果遇到的是奇数，就表示这个元素已经调整完成了，继续从左往右遍历，直到遇到一个偶数。
然后从nums 右侧开始遍历，如果遇到的是偶数，就表示这个元素已经调整完成了，继续从右往左遍历，直到遇到一个奇数。
交换这个偶数和奇数的位置，并且重复两边的遍历，直到在中间相遇，nums 调整完成。


链接：https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/solution/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-en35/
*/

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left]%2 == 1 {
			left += 1
		}
		for left < right && nums[right]%2 == 0 {
			right -= 1
		}
		nums[left], nums[right] = nums[right], nums[left]
		left += 1
		right -= 1
	}
	return nums
}

func main() {
	test := []int{3, 2, 4, 5}
	fmt.Println(exchange(test))
}
