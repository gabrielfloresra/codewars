package main

import (
	"fmt"
)

func rgb(r int, g int, b int, output *string) {

	var wrap = func(val int) int {
		if val > 255 {
			return 255
		} else if val < 0 {
			return 0
		}
		return val
	}

	*output = fmt.Sprintf("%02X%02X%02X", wrap(r), wrap(g), wrap(b))
}

func main() {

	var output string

	rgb(148, 0, 211, &output)
	fmt.Println(output)

	rgb(255, 255, 255, &output)
	fmt.Println(output)

	rgb(0, 0, 0, &output)
	fmt.Println(output)

	rgb(1, 2, 3, &output)
	fmt.Println(output)
}
