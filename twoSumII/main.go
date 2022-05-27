package main

func main() {
	twoSum([]int{5, 25, 75}, 100)

}

func twoSum(numbers []int, target int) []int {
	ln := len(numbers)
	var res []int
	for i := 0; i < ln-1; i++ {
		pivot := i + 1
		for pivot < ln {
			if numbers[i]+numbers[pivot] == target {
				res = append(res, i+1)
				res = append(res, pivot+1)
				return res
			}
			pivot++
		}
	}

	return res
}
