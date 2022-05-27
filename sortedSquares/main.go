package main

import (
	"fmt"
	"sort"
)

func main() {
	res := sortedSquares([]int{-4, -1, 0, 3, 10})
	fmt.Println(res)

}

func sortedSquares(nums []int) []int {
	for i, v := range nums {
		nums[i] = v * v
	}
	sort.Ints(nums)

	return nums
}
