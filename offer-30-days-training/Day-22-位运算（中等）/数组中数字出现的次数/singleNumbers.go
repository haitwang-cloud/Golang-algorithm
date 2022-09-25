package main

import "fmt"

//数组中数字出现的次数

//链接：https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/solution/jian-zhi-offer-56-i-shu-zu-zhong-shu-zi-tykom/
func singleNumbers(nums []int) []int {
	var res int
	//# 1. 遍历异或
	for _, v := range nums {
		res ^= v
	}
	//2. 循环左移，计算 m
	m := 1
	for res&m == 0 {
		m <<= 1
	}
	a, b := 0, 0
	//3. 遍历 nums 分组
	for _, v := range nums {
		//当 num & m == 0
		if v&m == 0 {
			a ^= v
		} else {
			//当 num & m != 0
			b ^= v
		}
	}
	return []int{a, b}
}

func main() {

	test := []int{4, 1, 4, 6}
	fmt.Println(singleNumbers(test))
}
