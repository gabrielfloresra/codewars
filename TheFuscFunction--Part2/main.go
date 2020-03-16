package main

import "math/big"

//Fusc function
func Fusc(n *big.Int) int {
	newN := big.NewInt(0).Div(n, big.NewInt(2))
	condicion := big.NewInt(0).DivMod()
	if n < 2 {
		return n
	} else if condicion {
		return Fusc(newN)
	}
	return Fusc(newN) + Fusc(newN+1)
}

func main() {

}
