package main

import "fmt"

func replaceSpace(s string) string {
	result := make([]byte, 0)
	strByte := []byte(s)
	for _, value := range strByte {
		if value == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, value)
		}
	}
	return string(result)
}

func main() {
	test := "We are happy."
	fmt.Println(replaceSpace(test))
}
