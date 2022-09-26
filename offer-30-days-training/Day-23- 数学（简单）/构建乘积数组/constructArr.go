package main

import "fmt"

//https://leetcode.cn/problems/gou-jian-cheng-ji-shu-zu-lcof/solution/mian-shi-ti-66-gou-jian-cheng-ji-shu-zu-biao-ge-fe/
func constructArr(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	b, tmp := make([]int, len(a)), 1
	b[0] = 1
	for index := 1; index < len(a); index++ { //计算下三角
		b[index] = b[index-1] * a[index-1]
	}
	for index := len(a) - 2; index >= 0; index-- { //计算上三角
		tmp *= a[index+1]
		b[index] *= tmp //下三角 * 上三角
	}
	return b
}

func main() {

	test := []int{1, 2, 3, 4, 5}
	fmt.Println(constructArr(test))
}
