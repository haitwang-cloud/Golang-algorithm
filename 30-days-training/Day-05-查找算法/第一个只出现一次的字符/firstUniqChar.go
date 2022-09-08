package main

import "fmt"

// hashtable 统计
func firstUniqChar(s string) byte {
	strDict := [26]int{}
	for _, v := range s {
		strDict[v-'a'] += 1
	}
	for index, value := range s {
		if strDict[value-'a'] == 1 {
			return s[index]
		}
	}
	return ' '
}

func main() {
	test := "abaccdeff"
	fmt.Println(firstUniqChar(string(test)))
}
