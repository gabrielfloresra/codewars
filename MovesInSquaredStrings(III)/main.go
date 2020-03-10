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
func rot90Clock(s string) string {
	var result string

	mat := stringToMatriz(s)
	matTemp := stringToMatriz(s)

	for i, val := range mat {
		n := len(val) - i - 1
		m := 0
		for j := range val {
			mat[m][n] = matTemp[i][j]
			m++
		}
		n--
	}

	for _, val := range mat {
		result += string([]rune(val)) + "\n"
	}

	return result[:len(result)-1]
}
func diag1Sym(s string) string {
	var result string

	mat := stringToMatriz(s)
	matTemp := stringToMatriz(s)

	for i, val := range mat {
		for j := range val {
			mat[i][j], mat[j][i] = matTemp[j][i], matTemp[i][j]
		}
	}

	for _, val := range mat {
		result += string([]rune(val)) + "\n"
	}

	return result[:len(result)-1]
}
func selfieAndDiag1(s string) string {
	var result string
	temp1 := strings.Split(s, "\n")
	temp2 := strings.Split(diag1Sym(s), "\n")

	for pos := range temp1 {
		result += temp1[pos] + "|" + temp2[pos] + "\n"
	}

	return result
}

type fParam func(string) string

func oper(f fParam, x string) string { return f(x) }

func main() {
	fmt.Println(oper(diag1Sym, "wuUyPC\neNHWxw\nehifmi\ntBTlFI\nvWNpdv\nIFkGjZ"))
	fmt.Println()
	fmt.Println(oper(rot90Clock, "rgavce\nvGcEKl\ndChZVW\nxNWgXR\niJBYDO\nSdmEKb"))
}
