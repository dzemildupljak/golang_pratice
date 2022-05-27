package main

import "fmt"

func main() {
	res := orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})
	// res := orangesRotting([][]int{{0, 2}})
	// res := orangesRotting([][]int{
	// 	{2, 1, 1},
	// 	{0, 1, 1},
	// 	{1, 0, 1}})
	// res := orangesRotting([][]int{
	// 	{2, 0, 1, 1, 1, 1, 1, 1, 1, 1},
	// 	{1, 0, 1, 0, 0, 0, 0, 0, 0, 1},
	// 	{1, 0, 1, 0, 1, 1, 1, 1, 0, 1},
	// 	{1, 0, 1, 0, 1, 0, 0, 1, 0, 1},
	// 	{1, 0, 1, 0, 1, 0, 0, 1, 0, 1},
	// 	{1, 0, 1, 0, 1, 1, 0, 1, 0, 1},
	// 	{1, 0, 1, 0, 0, 0, 0, 1, 0, 1},
	// 	{1, 0, 1, 1, 1, 1, 1, 1, 0, 1},
	// 	{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	// 	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}})

	fmt.Println(res)

}
func orangesRotting(grid [][]int) int {
	var queue [][2]int
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	count := 0
	for {
		q := len(queue)
		for i := 0; i < q; i++ {
			pop := queue[0]
			x, y := pop[0], pop[1]
			queue = queue[1:]

			d := [][]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
			for _, v := range d {
				if v[0] < 0 || v[1] < 0 || v[0] >= m || v[1] >= n {
					continue
				}
				if grid[v[0]][v[1]] == 1 {
					grid[v[0]][v[1]] = 2
					queue = append(queue, [2]int{v[0], v[1]})
				}
			}
		}
		if len(queue) == 0 {
			break
		}
		count++
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return count
}
