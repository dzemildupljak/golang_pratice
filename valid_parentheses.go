package main

import (
	"fmt"
)

func VPmain() {

	// res := ValidParentheses("(())((()())())") //true
	res := ValidParentheses("({}[])") //true

	fmt.Println(res)

}

func ValidParentheses(parens string) bool {
	// Your code goes here
	if len(parens)%2 != 0 {
		return false
	}

	arr := []string{}

	for _, v := range parens {
		if string(v) == "[" {
			arr = append(arr, "]")
			continue
		} else if string(v) == "(" {
			arr = append(arr, ")")
			continue
		} else if string(v) == "{" {
			arr = append(arr, "}")
			continue
		}

		if len(arr) == 0 || arr[len(arr)-1] != string(v) {
			return false
		}

		arr = arr[:len(arr)-1]
	}

	return len(arr) == 0
}
