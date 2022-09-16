package main

import "fmt"

/*

 	1. 双指针去掉首尾多余空格，同时转化为字节数组
	2. 翻转区间 [left, right] 左闭右闭区间
	3.翻转每个单词

*/

func reverseWords(s string) string {
	//1. 双指针去掉首尾多余空格，同时转化为字节数组
	trimSpaces := func(s string) []byte {
		left, right := 0, len(s)-1
		// 去掉两端空格
		for left <= right && s[left] == ' ' {
			left++
		}
		for left <= right && s[right] == ' ' {
			right--
		}
		// 去掉中间多余空格
		strByte := make([]byte, 0, right-left+1)
		for left <= right {
			if s[left] != ' ' {
				strByte = append(strByte, s[left]) // 不为空格则放入strByte
			} else if strByte[len(strByte)-1] != ' ' { // strByte最后一个字符不为空格则放入，此处保证了单词直接只保留1个空格
				strByte = append(strByte, s[left])
			}
			left++
		}
		return strByte
	}

	// 2. 翻转区间 [left, right] 左闭右闭区间
	reverse := func(strByte []byte, left, right int) {
		for ; left < right; left, right = left+1, right-1 {
			strByte[left], strByte[right] = strByte[right], strByte[left]
		}
	}

	// 3. 翻转每个单词
	reverseEachWord := func(strByte []byte) {
		// start,end：单词的起始位置，n：字符串长度
		start, end, n := 0, 0, len(strByte)
		for start < n {
			for end < n && strByte[end] != ' ' {
				end++
			}
			// 此时 strByte[start, end) 是一个单词，翻转之，注意：此时end指向的是空格的位置
			reverse(strByte, start, end-1)
			// 更新 start,end 去找下一个单词的起始位置
			start, end = end+1, end+1
		}
	}

	// main process
	strByte := trimSpaces(s)
	reverse(strByte, 0, len(strByte)-1)
	reverseEachWord(strByte)
	return string(strByte)

}

func main() {
	s := "   the sky is blue  "
	fmt.Println(reverseWords(s))
}
