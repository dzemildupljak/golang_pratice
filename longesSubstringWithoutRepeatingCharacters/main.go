package main

import (
	"fmt"
)

func main() {
	res := lengthOfLongestSubstring("atqytslqaaaa")
	// res := lengthOfLongestSubstring(" ")
	fmt.Println(res)

}

func lengthOfLongestSubstring(s string) int {
	chars := [128]bool{}
	ln, max := 0, 0
	for i, j := 0, 0; i < len(s); i++ {
		index := s[i]
		if chars[index] {
			for ; chars[index]; j++ {
				ln--
				chars[s[j]] = false
			}
		}
		chars[index] = true
		ln++
		if ln > max {
			max = ln
		}
	}

	return max
}

// func lengthOfLongestSubstring(s string) int {
// 	dict := [128]bool{}
// 	length, max := 0, 0
// 	for i, j := 0, 0; i < len(s); i++ {
// 		index := s[i]
// 		if dict[index] {
// 			for ; dict[index]; j++ {
// 				length--
// 				dict[s[j]] = false
// 			}
// 		}

// 		dict[index] = true
// 		length++
// 		if length > max {
// 			max = length
// 		}
// 	}
// 	return max
// }

// func lengthOfLongestSubstring(s string) int {
// 	m := make(map[rune]int)
// 	start, max := -1, 0

// 	for i, v := range s {
// 		last := m[v]
// 		if last > start {
// 			start = last
// 		}
// 		m[v] = i
// 		if i-start > max {
// 			max = i - start
// 		}
// 	}
// 	return max
// }
