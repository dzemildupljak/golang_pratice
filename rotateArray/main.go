package main

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
func rotate(nums []int, k int) {
	k = k % len(nums)
	rev(nums)
	rev(nums[:k])
	rev(nums[k:])
}

func rev(x []int) {
	n := len(x)
	for i := 0; i < n/2; i++ {
		x[i], x[n-1-i] = x[n-1-i], x[i]
	}
}

// func rotate(nums []int, k int) {
// 	r := k % len(nums)
// 	fmt.Println(r)
// 	for i := 0; i < k; i++ {
// 		tmp := nums[len(nums)-1]
// 		for j := len(nums) - 1; j > 0; j-- {
// 			nums[j] = nums[j-1]
// 		}
// 		nums[0] = tmp
// 	}

// 	fmt.Println(nums)

// }
