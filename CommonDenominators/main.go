package main

import (
	"strconv"
)

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a int, b int) int {
	return a / gcd(a, b) * b
}

func convertFracts(a [][]int) string {
	if a == nil || len(a) == 0 {
		return "(,)"
	}
	min := a[0][1] / gcd(a[0][0], a[0][1])
	for i := 1; i < len(a); i++ {
		min = lcm(min, a[i][1]/gcd(a[i][0], a[i][1]))
	}

	var cad string
	for _, val := range a {
		cad += "(" + strconv.Itoa((min*val[0])/val[1]) + "," + strconv.Itoa(min) + ")"
	}

	return cad
}
func main() {
	convertFracts([][]int{{69, 130}, {87, 1310}, {30, 40}})
}
