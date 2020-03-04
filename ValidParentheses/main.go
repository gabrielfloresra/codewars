package main

import (
	"fmt"
	"regexp"
)

func validParentheses(parens string) bool {

	// num := 0

	// for c := range parens {
	// 	if parens[c] == 40 {
	// 		num++
	// 	} else if parens[c] == 41 {
	// 		num--
	// 	}
	// 	if num < 0 {
	// 		return false
	// 	}
	// }

	// if num == 0 {
	// 	return true
	// }
	// return false
	_, err := regexp.Compile(parens)
	return err == nil
}

func main() {
	fmt.Println(validParentheses("()()()"))
}
