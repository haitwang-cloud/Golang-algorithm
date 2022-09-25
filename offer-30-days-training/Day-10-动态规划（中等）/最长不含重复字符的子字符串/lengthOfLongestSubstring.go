package main

import "fmt"

/*链接：https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/solution/jian-zhi-offer48zui-chang-bu-han-zhong-f-bpd2/

golang用到map的时候需要注意，如果不做ok的判断，字典中没有找到返回值也是0

会对第一个字母s[0]的判断有影响，所以这里用到 if closeIndex, ok := dict[val]; !ok 判断字典中是否存在

这里hashmap中记录的是每个子母在s中离当前最近的位置


*/
func lengthOfLongestSubstring(s string) int {
	/*

	 */
	strArray := []byte(s)
	strDict := map[byte]int{}
	dp := make([]int, len(strArray))
	maxLen := 0
	for index, val := range strArray {
		if index == 0 {
			// 当 index == 0时，即s[j]左边无相同字符,则dp[j]=dp[j-1]+1
			dp[index] = 1
			strDict[val] = 0
		} else if closeIndex, ok := strDict[val]; !ok {
			// 当 dp[j-1]<j-i时，即字符s[i]在子字符串dp[j-1]区间之外,则dp[j]=dp[j-1]+1
			dp[index] = dp[index-1] + 1
		} else if dp[index-1] < index-closeIndex {
			// 当 dp[j-1]<j-i时，即字符s[i]在子字符串dp[j-1]区间之外,则dp[j]=dp[j-1]+1
			dp[index] = dp[index-1] + 1
		} else {
			// 当 dp[j-1]>=j-i时，即字符s[i]在子字符串dp[j-1]区间之中,则dp[j]的左边界由s[i]决定,dp[j]=j-i
			dp[index] = index - closeIndex
		}
		strDict[val] = index
		maxLen = max(dp[index], maxLen)
	}
	return maxLen
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	test := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(test))
}
