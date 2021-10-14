package main

import "fmt"

func firstUniqChar(s string) byte {
	strDict := [26]int{}
	for _, value := range s {
		strDict[value-'a']++
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
	fmt.Println(firstUniqChar(test))
}
