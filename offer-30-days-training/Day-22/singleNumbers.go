package main

import "fmt"

//数组中数字出现的次数

func singleNumbers(nums []int) []int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	div := 1
	for res&div == 0 {
		div <<= 1
	}
	a, b := 0, 0
	for _, v := range nums {
		if v&div == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}

/*

    def singleNumbers(self, nums: List[int]) -> List[int]:
        x, y, n, m = 0, 0, 0, 1
        for num in nums:         # 1. 遍历异或
            n ^= num
        while n & m == 0:        # 2. 循环左移，计算 m
            m <<= 1
        for num in nums:         # 3. 遍历 nums 分组
            if num & m: x ^= num # 4. 当 num & m != 0
            else: y ^= num       # 4. 当 num & m == 0
        return x, y              # 5. 返回出现一次的数字

链接：https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/solution/jian-zhi-offer-56-i-shu-zu-zhong-shu-zi-tykom/
*/

func main() {

	test := []int{4, 1, 4, 6}
	fmt.Println(singleNumbers(test))
}
