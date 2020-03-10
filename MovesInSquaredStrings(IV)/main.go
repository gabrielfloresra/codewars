package main

import (
	"fmt"
	"strings"
)

func stringToMatriz(cad string) [][]rune {
	mat := [][]rune{}

	for _, cadena := range strings.Split(cad, "\n") {
		mat = append(mat, []rune(cadena))
	}
	return mat
}

//Rot90Counter function
func Rot90Counter(s string) string {
	var result string

	mat := stringToMatriz(s)
	matTemp := stringToMatriz(s)

	for i, val := range mat {
		m := len(val) - 1
		for j := range val {
			mat[m][i] = matTemp[i][j]
			m--
		}
	}

	for _, val := range mat {
		result += string([]rune(val)) + "\n"
	}

	return result[:len(result)-1]
}

//Diag2Sym function
func Diag2Sym(s string) string {
	var result string

	mat := stringToMatriz(s)
	matTemp := stringToMatriz(s)

	for i, val := range mat {
		m := len(val) - i - 1
		n := len(val) - 1
		for j := range val {
			mat[m][n] = matTemp[j][i]
			n--
		}
	}

	for _, val := range mat {
		result += string([]rune(val)) + "\n"
	}

	return result[:len(result)-1]
}

//SelfieDiag2Counterclock function
func SelfieDiag2Counterclock(s string) string {
	var result string

	temp1 := strings.Split(s, "\n")
	temp2 := strings.Split(Diag2Sym(s), "\n")
	temp3 := strings.Split(Rot90Counter(s), "\n")

	for pos := range temp1 {
		result += temp1[pos] + "|" + temp2[pos] + "|" + temp3[pos] + "\n"
	}
	return result[:len(result)-1]
}

//FParam type
type FParam func(string) string

//Oper function
func Oper(f FParam, x string) string { return f(x) }

func main() {
	fmt.Println(Oper(Rot90Counter, "abcd\nefgh\nijkl\nmnop"))
}
