package main

import "fmt"

//Fusc function
func Fusc(n int) int {

	condicion := true
	for condicion {
		if n < 2 {
			condicion = false
		}
		n = n/(2-(n%2)) + ((n / ((n * ((n % 2) - 1)) + (1 * (n % 2)))) + 1)
	}
	return n
}

//newN := int(n / 2)
//condicion := n%2 == 0
//if n < 2 {
//	return n
//} else if condicion {
//	return Fusc(newN)
//}
//return Fusc(newN) + Fusc(newN+1)
func main() {
	fmt.Println(Fusc(3))
}
