package main

import "fmt"

func numWays(n int) int {
	const mod int = 1e9 + 7
	nums1, nums2 := 1, 1
	for i := 2; i < n; i++ {
		nums1, nums2 = nums2, (nums1+nums2)%mod

	}
	return nums1
}

func main() {
	fmt.Println(numWays(7))
	fmt.Println(numWays(5))
	fmt.Println(numWays(1))

}
