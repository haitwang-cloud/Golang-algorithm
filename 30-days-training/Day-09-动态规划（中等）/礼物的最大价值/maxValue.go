package main

import "fmt"

/*
	f(i,j)=max(f(i-1,j),f(i,j-1))+ grid[i][j]
*/
func maxValue(grid [][]int) int {
	for height := 0; height < len(grid); height++ {
		for width := 0; width < len(grid[0]); width++ {
			if height == 0 && width == 0 {
				continue
			} else if height == 0 {
				grid[height][width] += grid[height][width-1]
			} else if width == 0 {
				grid[height][width] += grid[height-1][width]
			} else {
				grid[height][width] += max(grid[height-1][width], grid[height][width-1])
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//for i in range(len(grid)):
//for j in range(len(grid[0])):
//if i == 0 and j == 0: continue
//if i == 0: grid[i][j] += grid[i][j - 1]
//elif j == 0: grid[i][j] += grid[i - 1][j]
//else: grid[i][j] += max(grid[i][j - 1], grid[i - 1][j])
//return grid[-1][-1]

func main() {
	test := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1}}
	fmt.Println(maxValue(test))
}
