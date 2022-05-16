package main

import "fmt"

/*

https://leetcode.cn/problems/partition-labels/solution/hua-fen-zi-mu-qu-jian-by-leetcode-solution/

*/

func partitionLabels(s string) (partition []int) {
	lastPos := [26]int{}
	for index, str := range s {
		lastPos[str-'a'] = index
	}
	start, end := 0, 0
	for index, char := range s {
		if lastPos[char-'a'] > end {
			end = lastPos[char-'a']
		}
		if index == end {
			partition = append(partition, end-start+1)
			start = end + 1

		}

	}
	return

}

func main() {
	test := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(test))
}
