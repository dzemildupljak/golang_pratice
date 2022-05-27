package main

import "fmt"

func main() {
	moveZeroes([]int{0, 1, 0, 3, 12})
	moveZeroes([]int{0, 0, 1})

}

func moveZeroes(nums []int) {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	fmt.Println(nums)
}

// func moveZeroes(nums []int) {
// 	ln := len(nums) - 1
// 	for i := 0; i < ln; i++ {
// 		if i > 0 && nums[i-1] == 0 {
// 			i = i - 1
// 		}
// 		if nums[i] == 0 {
// 			for j := i; j < ln; j++ {
// 				nums[j], nums[j+1] = nums[j+1], nums[j]
// 			}
// 		}
// 	}
// }
