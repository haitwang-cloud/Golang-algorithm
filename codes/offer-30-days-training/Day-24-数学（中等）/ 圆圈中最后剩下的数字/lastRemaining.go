package main

import "fmt"

//链接：https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/solution/huan-ge-jiao-du-ju-li-jie-jue-yue-se-fu-huan-by-as/
func lastRemaining(n int, m int) int {

	pos := 0 // 最终活下来那个人的初始位置
	for index := 2; index <= n; index++ {
		pos = (pos + m) % index // 每次循环右移
	}
	return pos
}
func main() {
	n, m := 5, 3
	fmt.Println(lastRemaining(n, m))
}
