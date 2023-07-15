package permute

// https://leetcode.cn/problems/permutations/
/*
	1.permute 函数初始化一个空的二维切片 res，用于存储所有的排列结果。
	2.它调用 backtrack 函数，传入 nums、起始索引 0，以及 res 的指针。
	backtrack 函数是排列生成算法的核心。
	3.递归的基本情况是当起始索引 start 等于 nums 切片的长度时。这意味着所有元素都已经固定在位置上，一个排列完成。
	4.在基本情况下，创建一个临时切片 temp，长度与 nums 相同，并使用 copy 函数将当前的 nums 状态复制到 temp 中。然后，将 temp 添加到 res 切片中。
	5.如果没有达到基本情况，函数进入一个循环，从当前的起始索引 start 开始，并迭代到 nums 切片的末尾。
	6.在循环内部，使用并行赋值将 nums 切片中索引 start 和 i 处的元素进行交换。
	7.函数对自身进行递归调用，将起始索引增加 1，以探索决策树的下一层。
	8.递归调用返回后，撤销之前的交换，以回溯并探索其他可能性。
	9.循环完成后，backtrack 函数退出，permute 函数返回存储在 res 中的生成排列结果。
*/

func permute(nums []int) [][]int {
	res := [][]int{}
	backtrack(nums, 0, &res)
	return res
}

// 回溯函数
func backtrack(nums []int, start int, res *[][]int) {
	// 触发结束条件
	if start == len(nums) {
		// 因为 nums 是全局变量，因此不需要创建新的数组，直接将 nums 添加到结果集
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
		return
	}
	for i := start; i < len(nums); i++ {
		// 做选择
		/*
		   通过交换数组元素的方式来避免使用额外的布尔数组来保证不会使用重复的数字。
		   具体做法是，每次固定一个位置的元素，
		   然后交换该位置与当前位置的元素，这样就可以避免重复使用相同的数字。
		*/
		nums[start], nums[i] = nums[i], nums[start]
		// 进入下一层决策树
		backtrack(nums, start+1, res)
		// 撤销选择
		nums[start], nums[i] = nums[i], nums[start]
	}
}
