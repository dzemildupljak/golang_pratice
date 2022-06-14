package main

import "fmt"

func main() {

	nums := []int{5, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8}
	// nums := []int{1}
	target := 8
	res := searchRange(nums, target)
	fmt.Println(res)

}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	left, right := 0, len(nums)-1
	for left+1 < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			left = mid
			right = mid
			for {
				left--
				if left < 0 || nums[left] != target {
					left++
					break
				}
			}
			for {
				right++
				if right > len(nums)-1 || nums[right] != target {
					right--
					break
				}
			}
			return []int{left, right}
		} else if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}

	if nums[left] == target && nums[right] == target {
		return []int{left, right}
	}
	if nums[left] == target {
		return []int{left, left}
	}
	if nums[right] == target {
		return []int{right, right}
	}
	return []int{-1, -1}
}

// func searchRange(nums []int, target int) []int {
// 	if len(nums) == 0 || target > nums[len(nums)-1] {
// 		return []int{-1, -1}
// 	}

// 	targetIndex := binarySearch(nums, 0, len(nums), target)
// 	if targetIndex == -1 {
// 		return []int{-1, -1}
// 	}

// 	left := targetIndex
// 	right := targetIndex

// 	for {
// 		if left-1 >= 0 && nums[left-1] == target {
// 			left -= 1
// 		}
// 		if right+1 < len(nums) && nums[right+1] == target {
// 			right += 1
// 		}

// 		if (right+1 > len(nums)-1 || nums[right+1] != target) &&
// 			(left-1 == -1 || nums[left-1] != target) {
// 			break
// 		}
// 	}
// 	return []int{left, right}
// }

// func binarySearch(numbers []int, leftBound, rightBound, numberToFind int) int {
// 	if rightBound >= leftBound {
// 		midPoint := leftBound + (rightBound-leftBound)/2

// 		if numbers[midPoint] == numberToFind {
// 			return midPoint
// 		}

// 		if numbers[midPoint] > numberToFind {
// 			return binarySearch(numbers, leftBound, midPoint-1, numberToFind)
// 		}

// 		return binarySearch(numbers, midPoint+1, rightBound, numberToFind)
// 	}

// 	return -1
// }
