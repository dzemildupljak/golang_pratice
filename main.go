package main

import "fmt"

func main() {
	nQueens(5)
}

func nQueens(n int) [][]string {
	var matrix [][]string

	for i := 0; i < n; i++ {
		row := make([]string, 0)
		for j := 0; j < n; j++ {

		}
		matrix = append(matrix, row)
	}

	fmt.Println(matrix)
	return matrix
}
