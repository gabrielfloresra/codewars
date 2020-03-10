package main

import (
	"fmt"
	"math"
	"strings"
)

func setTLenght(s string, num int) string {
	for len(s) < num {
		s += "\v"
	}
	return s
}

func removeTabs(s string) string {
	for i := len(s) - 1; s[i] == '\v'; i-- {
		s = s[:i]
	}
	return s
}

func getNValue(lNum int) int {

	sqrt := math.Sqrt(float64(lNum))
	sqrtNum := math.Round(sqrt)

	if sqrtNum < sqrt {
		return int(sqrtNum) + 1
	}
	return int(sqrtNum)
}

func stringToMatriz(cad string) [][]rune {
	mat := [][]rune{}
	for _, cadena := range strings.Split(cad, "\n") {
		mat = append(mat, []rune(cadena))
	}
	return mat
}

//Rot90Clock function
func Rot90Clock(s string) string {
	var result string
	mat := stringToMatriz(s)
	matTemp := stringToMatriz(s)

	for i, val := range mat {
		n := len(val) - i - 1
		m := 0
		for j := range val {
			matTemp[m][n] = mat[i][j]
			m++
		}
		n--
	}

	for _, val := range matTemp {
		result += string([]rune(val)) + "\n"
	}

	return result[:len(result)-1]
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

//Code funtion
func Code(s string) string {
	//s = strings.Replace(s, " ", ".", -1)}
	if len(s) == 0 {
		return ""
	}
	n := getNValue(len(s))
	s = setTLenght(s, n*n)
	var sNew string
	for pos, val := range s {
		if pos%n == 0 {
			sNew += "\n"
		}
		sNew += string(val)
	}
	return Rot90Clock(sNew[1:])
}

//Decode funtion
func Decode(s string) string {
	sNew := Rot90Counter(s)
	//sNew = strings.Replace(sNew, ".", " ", -1)
	return removeTabs(strings.Join(strings.Split(sNew, "\n"), ""))
}

func main() {
	//s := Code("Process terminated with status 0 (0 minute(s), 6 second(s))")
	s := Code("I.was.going.fishing.that.morning.at.ten.o'clock\nProcess terminated with status 0 (0 minute(s), 6 second(s))\nWhat do you remember? When I looked at his streaky glasses, I wanted to leave him. And before that? He stole those cherries for me at midnight. We were walking in the rain and I loved him. And before that? I saw him coming toward me that time at the picnic, edgy, foreign.\nSky a shook poncho. Roof wrung. Mind a luna moth in a banjo.This weather's witty Peek-a-boo. A study in Insincerity\nSome say the world will end in fire, Some say in ice. From what I've tasted of desire I hold with those who favor fire. But if it had to perish twice, I think I know enough of hate To say that for destruction ice Is also great And would suffice.\nThe panel found that he was loyal and discreet with atomic secrets, but did not recommend that his security clearance be reinstated. This ended his role in government and policymaking. He became an academic exile, cut off from his former career and the world he had helped to create.")
	fmt.Println(s)
	fmt.Printf(Decode(s))
}
