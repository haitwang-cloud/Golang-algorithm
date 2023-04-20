package main

func fib(n int) int {
	if n < 2 {
		return n
	}
	/* 递归
	递归的时间复杂度是指数级别的，这里的时间复杂度是 O(2^n)
	return fib(n-1) + fib(n-2)
	*/
	/* 动态规划
	动态规划的时间复杂度是线性级别的，这里的时间复杂度是 O(n)

	memo := make([]int, n+1)
	memo[0], memo[1] = 0, 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
	*/
	nums1, nums2 := 0, 1
	for i := 2; i <= n; i++ {
		nums1, nums2 = nums2, nums1+nums2
	}
	return nums2
}
