package main

import "fmt"

/*
	f(n)=f(n-1)+f(n-2)
	青蛙跳台阶问题,f(0)=1,f(1)=1
	斐波那契问题，f(0)=0,f(1)=1
*/

func numWays(n int) int {
	const mod int = 1e9 + 7
	nums1, nums2 := 1, 1
	for index := 0; index < n; index++ {
		nums1, nums2 = nums2, (nums1+nums2)%mod
	}
	return nums1
}

func main() {
	fmt.Println(numWays(7))
	fmt.Println(numWays(2))
	fmt.Println(numWays(1))

}
