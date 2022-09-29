package main

import (
	"math"
	"strings"
)

//链接：https://leetcode.cn/problems/ba-zi-fu-chuan-zhuan-huan-cheng-zheng-shu-lcof/solution/mian-shi-ti-67-ba-zi-fu-chuan-zhuan-huan-cheng-z-4/

func strToInt(str string) int {

	str = strings.TrimSpace(str)
	result := 0
	sign := 1

	for i, v := range str {
		if v >= '0' && v <= '9' {
			result = result*10 + int(v-'0')
		} else if v == '-' && i == 0 {
			sign = -1
		} else if v == '+' && i == 0 {
			sign = 1
		} else {
			break
		}

		if result > math.MaxInt32 {
			if sign == -1 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}

	return sign * result

}

func main() {
	test := []string{"-+12", "42", "   -42", "4193 with words", "words and 987", "-91283472332", "3.14159", "  0000000000012345678"}
	for _, v := range test {
		println(strToInt(v))
	}
}
