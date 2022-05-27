package main

func main() {
	containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})
}

func containsDuplicate(nums []int) bool {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		dict[nums[i]]++
		if dict[nums[i]] > 1 {
			return true
		}
	}

	return false
}
