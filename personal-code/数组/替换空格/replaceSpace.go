package main

import "fmt"

func replaceSpace(s string) string {
	ans := ""
	for _,v := range s{
		switch v{
		case ' ':
			ans = ans + "%20"
		default:
			ans = ans + string(v)
		}
	}
	return ans
}

func main() {
	test := "We are happy."
	fmt.Println(replaceSpace(test))
}
