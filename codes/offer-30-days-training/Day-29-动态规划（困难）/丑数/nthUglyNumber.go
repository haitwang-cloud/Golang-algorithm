package main

import "fmt"

//https://leetcode.cn/problems/chou-shu-lcof/solution/mian-shi-ti-49-chou-shu-dong-tai-gui-hua-qing-xi-t/

func nthUglyNumber(n int) int {

	dp := make([]int, n)
	for index := 0; index < n; index++ {
		dp[index] = 1
	}
	a, b, c := 0, 0, 0
	for index := 1; index < n; index++ {
		n2, n3, n5 := dp[a]*2, dp[b]*3, dp[c]*5
		dp[index] = min(min(n2, n3), n5)
		if dp[index] == n2 {
			a++
		}
		if dp[index] == n3 {
			b++
		}
		if dp[index] == n5 {
			c++
		}

	}
	return dp[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/*

   def nthUglyNumber(self, n: int) -> int:
       dp, a, b, c = [1] * n, 0, 0, 0
       for i in range(1, n):
           n2, n3, n5 = dp[a] * 2, dp[b] * 3, dp[c] * 5
           dp[i] = min(n2, n3, n5)
           if dp[i] == n2: a += 1
           if dp[i] == n3: b += 1
           if dp[i] == n5: c += 1
       return dp[-1]
*/

func main() {
	test := 10
	fmt.Println(nthUglyNumber(test))

}
