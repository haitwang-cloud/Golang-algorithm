package main

import (
	"fmt"
)

//https://leetcode.cn/problems/fruit-into-baskets/

func totalFruit(fruits []int) int {

	//建立一个哈希表，记录已经在篮子里的水果个数(区间内元素的种类个数)
	counter := make(map[int]int)

	start, end, result := 0, 0, 0
	for end < len(fruits) {
		///我们想得到以end为结尾的区间，所以end必须取到
		counter[fruits[end]]++
		//当哈希表的长度是3的时候表示区间内元素的种类有3个，不符合题意
		for start <= end && len(counter) >= 3 {
			//start指针向右移动，将start指向的水果数目-1
			counter[fruits[start]]--
			//当hashMap[fruits[start]] == 0时表示篮子里没有这种水果了，直接delete
			if counter[fruits[start]] == 0 {
				delete(counter, fruits[start])
			}
			start++
		}

		//每次循环之后，我们都得到以end为结尾的符合题意的最长区间
		//通过if选择每次的结果，得到最终的最优解
		if end-start+1 > result {
			result = end - start + 1
		}
		end++
	}
	return result

}

func main() {
	fruits := []int{1, 2, 1}
	fmt.Println(totalFruit(fruits))
}
