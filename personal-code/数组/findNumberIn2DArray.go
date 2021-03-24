package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	width, height := len(matrix[0]), len(matrix)
	for h, w := 0, width-1; w >= 0 && h < height; {
		if matrix[h][w] == target {
			return true
		} else if matrix[h][w] > target {
			w--
		} else {
			h++
		}

	}
	return false
}

func main() {
	nums, target := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9
	nums1, target1 := [][]int{{-5}}, -5
	nums2, target2 := [][]int{{1, 1}}, 0
	fmt.Println(findNumberIn2DArray(nums, target))
	fmt.Println(findNumberIn2DArray(nums1, target1))
	fmt.Println(nums2[0][1])
	fmt.Println(findNumberIn2DArray(nums2, target2))

}
