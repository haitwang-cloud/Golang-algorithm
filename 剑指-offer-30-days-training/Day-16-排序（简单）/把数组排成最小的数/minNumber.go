package main

import (
	"fmt"
	"strconv"
)

//快排实现排序，先排序后转成string
//链接：https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/solution/kuai-pai-shi-xian-pai-xu-xian-pai-xu-hou-4sil/

func minNumber(nums []int) string {

	// 快排实现排序，排序后转成string
	res := make([]string, len(nums))
	for i, v := range nums {
		res[i] = strconv.Itoa(v)
	}
	compare := func(str1, str2 string) bool {
		num1, _ := strconv.Atoi(str1 + str2)
		num2, _ := strconv.Atoi(str2 + str1)
		if num1 < num2 {
			return true
		}
		return false
	}
	var quickSort func(strArr []string, l, r int)
	quickSort = func(strArr []string, l, r int) {
		if l >= r {
			return
		}
		i, j, pivot := l, r, strArr[l]
		for i < j {
			for i < j && compare(pivot, strArr[j]) {
				j--
			}
			for i < j && !compare(pivot, strArr[i]) {
				i++
			}
			strArr[i], strArr[j] = strArr[j], strArr[i]
		}
		strArr[i], strArr[l] = strArr[l], strArr[i]
		quickSort(strArr, l, i-1)
		quickSort(strArr, i+1, r)
	}
	quickSort(res, 0, len(nums)-1)
	ans := ""
	for _, s := range res {
		ans += s
	}
	return ans

}

func main() {
	test := []int{2, 7, 11, 15}
	fmt.Println(minNumber(test))
}
