package main

import "fmt"

/*

本问题是典型的矩阵搜索问题，可使用 深度优先搜索（DFS）+ 剪枝 解决。

深度优先搜索： 可以理解为暴力法遍历矩阵中所有字符串可能性。DFS 通过递归，先朝一个方向搜到底，再回溯至上个节点，沿另一个方向搜索，以此类推。
剪枝： 在搜索中，遇到 这条路不可能和目标字符串匹配成功 的情况（例如：此矩阵元素和目标字符不同、此元素已被访问），则应立即返回，称之为 可行性剪枝 。

链接：https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/solution/mian-shi-ti-12-ju-zhen-zhong-de-lu-jing-shen-du-yo/
*/

func exist(board [][]byte, word string) bool {
	height, width := len(board), len(board[0])
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			//如果在数组中找得到第一个数，就执行下一步，否则返回false
			if search(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}
func search(board [][]byte, i, j, k int, word string) bool {
	//如果找到最后一个数，则返回true,搜索成功
	if k == len(word) {
		return true
	}
	//i,j的约束条件
	//(1) 行或列索引越界
	//(2) 当前矩阵元素与目标字符不同
	//(3) 当前矩阵元素已访问过 （代表之前已经搜索过，不需要再搜索）
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return false
	}

	//进入DFS深度优先搜索
	//先把正在遍历的该值重新赋值，如果在该值的周围都搜索不到目标字符，则再把该值还原
	//如果在数组中找到第一个字符，则进入下一个字符的查找
	if board[i][j] == word[k] {
		temp := board[i][j]
		board[i][j] = ' '
		//下面这个if语句，如果成功进入，说明找到该字符，然后进行下一个字符的搜索,直到所有的搜索都成功，
		//即k == len(word) - 1 的大小时，会返回true，进入该条件语句，然后返回函数true值。
		if search(board, i, j+1, k+1, word) ||
			search(board, i, j-1, k+1, word) ||
			search(board, i+1, j, k+1, word) ||
			search(board, i-1, j, k+1, word) {
			return true
		} else {
			//没有找到目标字符，还原之前重新赋值的board[i][j]
			board[i][j] = temp
		}
	}

	return false
}

func main() {
	test, word := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"
	fmt.Println(exist(test, word))
}
