package main

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
		   通过交换数组元素的方式来避免使用额外的布尔数组。
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

/*

当你调用 `permute` 函数时，它将返回给定整数数组的所有全排列。

在 `permute` 函数内部，我们首先创建了一个空的结果集 `res`，用于存储所有的全排列。

接下来，我们定义了三个辅助函数：

1. `backtrack` 函数是一个递归函数，它通过回溯算法来生成全排列。它接受四个参数：原始数组 `nums`、当前已经生成的排列 `track`、一个布尔数组 `used` 用于标记已经使用过的数字、以及结果集 `res` 的指针。

2. `backtrack` 函数中的结束条件是 `len(track) == len(nums)`，即当已经生成的排列长度等于原始数组的长度时，说明已经生成了一个全排列。此时，我们创建一个临时数组 `temp` 来存储一份全排列，并将其添加到结果集 `res` 中。

3. 在 `backtrack` 函数的主体部分，我们使用一个循环遍历原始数组 `nums` 中的元素。在每次迭代中，我们进行以下操作：
   - 首先，我们检查当前位置的数字是否已经被使用，通过查看布尔数组 `used` 中对应位置的值来确定。如果已经使用过，我们就跳过该数字，进行下一次迭代（剪枝操作）。
   - 如果当前位置的数字没有被使用，我们就将其加入到当前已生成的排列 `track` 中，然后将布尔数组 `used` 中对应位置的值设置为 `true`，表示该数字已经被使用。
   - 接下来，我们递归调用 `backtrack` 函数，继续生成下一个位置的数字。
   - 在递归调用返回后，我们进行撤销选择的操作，即将当前位置的数字从 `track` 中移除，将布尔数组 `used` 中对应位置的值恢复为 `false`，以便在后续的迭代中可以选择该数字。

最后，在 `permute` 函数中，我们调用了 `backtrack` 函数来开始生成全排列，并将结果集 `res` 返回。

总结起来，该算法使用回溯算法来生成给定整数数组的全排列。通过遍历数组中的每个位置，进行选择和撤销选择的操作，递归地生成所有可能的排列，最终得到结果集。
*/
