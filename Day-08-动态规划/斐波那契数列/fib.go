package main

import "fmt"

func fib(n int) int {
	const mod int = 1e9 + 7
	if n < 0 {
		return -1
	} else if n <= 1 {
		return n
	}
	nums1, nums2, value := 0, 1, 1
	for i := 2; i < n; i++ {
		nums1 = nums2
		nums2 = value
		value = (nums1 + nums2) % mod

	}
	return value
}

func main() {
	fmt.Println(fib(5))
	fmt.Println(fib(1))

}
