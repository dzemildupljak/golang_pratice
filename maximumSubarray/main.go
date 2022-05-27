package main

import (
	"fmt"
	"math"
)

func main() {
	res := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	fmt.Println(res)
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sums := make([]int, len(nums))

	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = max(nums[i]+sums[i-1], nums[i])
	}

	maxVal := math.MinInt32

	for i := 0; i < len(sums); i++ {
		maxVal = max(sums[i], maxVal)
	}

	return maxVal
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
