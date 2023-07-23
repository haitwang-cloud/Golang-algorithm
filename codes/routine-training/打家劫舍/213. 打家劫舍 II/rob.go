package rob

// https://leetcode.cn/problems/house-robber-ii/
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

func robRange(nums []int, start, end int) int {

	dp_i_1, dp_i_2, dp_i := 0, 0, 0
	for i := end; i >= start; i-- {
		dp_i = max(dp_i_1, dp_i_2+nums[i])
		dp_i_2, dp_i_1 = dp_i_1, dp_i
	}
	return dp_i

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
