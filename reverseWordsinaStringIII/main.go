package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))

}
func reverseWords(s string) string {
	res := strings.Split(s, " ")
	fmt.Println(res)
	for i, v := range res {
		res[i] = string(reverseString([]byte(v)))
	}
	return strings.Join(res, " ")

}

func reverseString(s []byte) []byte {
	ln := len(s)
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
