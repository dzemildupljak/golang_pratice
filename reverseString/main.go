package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(s)

	reverseString(s)
	fmt.Println(s)
}

func reverseString(s []byte) {
	ln := len(s)
	for i, j := 0, ln-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
