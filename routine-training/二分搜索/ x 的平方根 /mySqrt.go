package main

import "fmt"

// https://leetcode.cn/problems/sqrtx/solution/x-de-ping-fang-gen-by-leetcode-solution/

/*

由于x 平方根的整数部分ans是满足k^2≤x的最大k值，因此我们可以对k进行二分查找，从而得到答案。
二分查找的下界为0，上界可以粗略地设定为x。在
二分查找的每一步中，我们只需要比较中间元素mid的平方与x的大小关系，并通过比较的结果调整上下界的范围
。由于我们所有的运算都是整数运算，不会存在误差，因此在得到最终的答案ans后，也就不需要再去尝试ans+1 了。

*/
func mySqrt(x int) int {
	if x < 2 {
		return x
	}
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		} else if mid*mid <= x {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return right
}

func main() {
	test := 9
	fmt.Println(mySqrt(test))
}
