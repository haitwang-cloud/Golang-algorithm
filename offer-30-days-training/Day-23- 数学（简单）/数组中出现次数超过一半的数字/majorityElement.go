package main

import "fmt"

//https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/solution/mian-shi-ti-39-shu-zu-zhong-chu-xian-ci-shu-chao-3/
func majorityElement(nums []int) int {
	vote, x, counter := 0, 0, 0
	for _, value := range nums {
		if vote == 0 {
			x = value
		}
		if value == x {
			vote++
		} else {
			vote--
		}

	}
	//验证 x 是否为众数
	for _, value := range nums {
		if value == x {
			counter++
		}
	}
	if counter*2 > len(nums) {
		return x
	}
	return 0
}
func main() {

	test := []int{1, 2, 2, 2, 5, 2}
	fmt.Println(majorityElement(test))
}
