package main

import "fmt"

//https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/solution/jian-zhi-offer-57-ii-he-wei-s-de-lian-xu-t85z/

func findContinuousSequence(target int) [][]int {
	i, j, sum := 1, 2, 3
	var res [][]int
	for i < j {
		if sum == target {
			var list []int
			for k := i; k <= j; k++ {
				list = append(list, k)
			}
			res = append(res, list)
			sum -= i
			i++
		} else if sum > target {
			sum -= i
			i++
		} else {
			j++
			sum += j
		}

	}
	return res
}

/*

class Solution:
    def findContinuousSequence(self, target: int) -> List[List[int]]:
        i, j, s, res = 1, 2, 3, []
        while i < j:
            if s == target:
                res.append(list(range(i, j + 1)))
            if s >= target:
                s -= i
                i += 1
            else:
                j += 1
                s += j
        return res


*/

func main() {
	test := 15
	fmt.Println(findContinuousSequence(test))

}
