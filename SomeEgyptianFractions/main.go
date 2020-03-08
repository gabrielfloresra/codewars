package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func mcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return mcd(b, a%b)
}

func decimalToRacional(frac string) (int, int) {

	decimales := frac[strings.Index(frac, ".")+1:]

	numerador, _ := strconv.Atoi(decimales)
	parteEntera, _ := strconv.Atoi(frac[:strings.Index(frac, ".")])
	denominador := int(math.Pow(10, float64(len(decimales))))
	numerador += parteEntera * denominador
	numMCD := mcd(numerador, int(denominador))

	return numerador / numMCD, denominador / numMCD
}

func decompose(frac string) []string {

	frac = strings.Replace(frac, " ", "", -1)
	if frac == "" {
		return []string{}
	}
	var numerador, denominador int
	var err error
	numeradorExist := false
	denominadorExist := false
	if strings.Contains(frac, ".") {
		numerador, denominador = decimalToRacional(frac)
	} else {

		data := strings.Split(frac, "/")
		numerador, err = strconv.Atoi(data[0])
		if err == nil {
			numeradorExist = true
		}
		if len(data) > 1 {
			denominador, err = strconv.Atoi(data[1])
			if err == nil {
				denominadorExist = true
			}
		} else {
			denominadorExist = false
		}
	}
	if !denominadorExist && numeradorExist && numerador != 0 {
		return []string{strconv.Itoa(numerador)}
	}
	if denominador == 0 || numerador == 0 {
		return []string{}
	}

	mcdND := mcd(numerador, denominador)
	numerador /= mcdND
	denominador /= mcdND
	num := 2
	var soluciones []string
	if numerador >= denominador {
		numVeces := numerador / denominador
		numerador -= denominador * numVeces
		soluciones = append(soluciones, strconv.Itoa(numVeces))
		fmt.Println(soluciones)
	}
	fmt.Println(soluciones)
	for numerador != 0 {
		if numerador < 0 {
			break
		}
		denominadorTemp := denominador
		numeradorTemp := numerador
		fmt.Printf("%d/%d - %d\n", numerador, denominador, num)
		if num%denominador != 0 {

			denominadorTemp *= num
			numeradorTemp *= num
		}
		if denominadorTemp/num <= numeradorTemp {
			numeradorTemp -= denominadorTemp / num
			soluciones = append(soluciones, "1/"+strconv.Itoa(num))
			numerador = numeradorTemp
			denominador = denominadorTemp
			mcdND := mcd(numerador, denominador)
			numerador /= mcdND
			denominador /= mcdND
			if numerador == 1 {
				numerador = 0
				soluciones = append(soluciones, "1/"+strconv.Itoa(denominador))
			}
		}
		num++
	}
	/*
		 mejor resultado de codewars (robado :v)
			for demon > 0 && numer > 0 {
				nextNum := demon/numer + 1
				if (nextNum-1)*numer == demon {
					nextNum -= 1
				}
				res = append(res, fmt.Sprintf("1/%d", nextNum))
				numer = numer*nextNum - demon
				demon *= nextNum
			}
	*/

	fmt.Printf("frac: %s\n", frac)
	fmt.Println(soluciones)
	return soluciones
}

func main() {
	decompose("50/4187")
	//fmt.Println(decimalToRacional("50/4187"))
}
