package main

import (
	"strings"
)

func ValidBraces(braces string) bool {
	braces = strings.ReplaceAll(braces, " ", "")
	for strings.Contains(braces, "()") || strings.Contains(braces, "[]") || strings.Contains(braces, "{}") {
		if strings.Contains(braces, "()") {
			braces = strings.ReplaceAll(braces, "()", "")
		} else if strings.Contains(braces, "[]") {
			braces = strings.ReplaceAll(braces, "[]", "")
		} else {
			braces = strings.ReplaceAll(braces, "{}", "")
		}
	}
	return len(braces) == 0
}

func main() {
	// fmt.Println(validBraces("(){}[]"))
	// fmt.Println(validBraces("([{}])"))
	// fmt.Println(validBraces("(}"))
	// fmt.Println(validBraces("[(])"))
	// fmt.Println(validBraces("[({})](]"))
}
