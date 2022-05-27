package main

import "fmt"

func main() {
	// fmt.Println(checkInclusion("ab", "eidbaooo"))
	// fmt.Println(checkInclusion("ab", "eidboaoo"))
	fmt.Println(checkInclusion("adc", "dcda"))

}
func checkInclusion(s1 string, s2 string) bool {
	// word := make(map[rune]int)
	// for i := 0; i < len(s1); i++ {
	// 	word[rune(s1[i])]++
	// }

	// for i := 0; i <= len(s2)-len(s1); i++ {
	// 	w := make(map[rune]int)
	// 	for i, v := range word {
	// 		w[i] = v
	// 	}
	// 	// chars := make(map[rune]int)
	// 	for j, k := i, 0; j < i+len(s1); j++ {
	// 		// chars[rune(s2[j])]++
	// 		w[rune(s2[j])]--
	// 		if w[rune(s2[j])] < 0 {
	// 			break
	// 		}
	// 		k++
	// 		if k == len(s1) {
	// 			return true
	// 		}
	// 	}
	// }

	// return false
	buf := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		buf[int(s1[i]-'a')]++
	}

	tmp := make([]int, 26)
	for i := 0; i+len(s1)-1 < len(s2); {
		copy(tmp, buf)
		j := i
		for ; j < i+len(s1); j++ {
			tmp[int(s2[j]-'a')]--
			if tmp[int(s2[j]-'a')] < 0 {
				break
			}
		}
		if j == i+len(s1) {
			return true
		}
		i++
	}
	return false
}
