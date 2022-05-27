package main

import "fmt"

func main() {
	res := maxAreaOfIsland(
		[][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}},
	)
	fmt.Println(res)
}

func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			res := 0
			if grid[i][j] == 1 {
				dfs(grid, i, j, &res)
			}
			if res > max {
				max = res
			}
		}
	}
	return max
}

func dfs(grid [][]int, r, c int, ct *int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] == 0 {
		return
	}
	grid[r][c] = 0
	*ct = *ct + 1
	dfs(grid, r-1, c, ct)
	dfs(grid, r+1, c, ct)
	dfs(grid, r, c-1, ct)
	dfs(grid, r, c+1, ct)
}
