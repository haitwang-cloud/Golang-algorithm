package main

import "fmt"

func replaceSpace(s string) string {
	result := ""
	for _, v := range s {
		switch v {
		case ' ':
			result = result + "%20"
		default:
			result = result + string(v)
		}

	}
	return result
}

func main() {
	test := "We are happy."
	fmt.Println(replaceSpace(test))
}
