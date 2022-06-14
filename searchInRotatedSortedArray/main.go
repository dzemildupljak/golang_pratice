package main

import "fmt"

func main() {
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	// nums := []int{12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	nums := []int{12, 13, 14, 17, 18, 19, 20, 26, 27, 28, 29, 4, 5, 6, 7, 8, 9, 10, 11}
	target := 29
	res := search(nums, target)
	fmt.Println(res)
}

func search(nums []int, target int) int {
	nLen := len(nums)
	if nLen == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	left := 0
	right := nLen - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			if nums[mid] < target || target < nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if nums[mid] > target || target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
