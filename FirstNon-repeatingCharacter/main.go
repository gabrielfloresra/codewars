package main

import (
	"fmt"
	"regexp"
)

func firstNonRepeatingLetter(cad string) (char string) {

	for len(cad) != 0 {
		charAux := string(cad[0])
		if charAux == "." {
			charAux = `\.`
		}
		rex, _ := regexp.Compile("(?i)" + charAux + "+")
		cadAux := rex.ReplaceAllString(cad, "")
		if len(cad)-1 == len(cadAux) {
			char = charAux
			break
		} else {
			cad = cadAux
		}
	}
	if char == `\.` {
		char = "."
	}
	return char
}

func main() {
	fmt.Println(firstNonRepeatingLetter("P.jdWNB6LPpQ;8a2TwAew:tTjLNwWc8rWNvsn"))
}
