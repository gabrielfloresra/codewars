package main

import(
	"fmt"
	"strings"
)

func spinWords(str string) string{
	var cadenaNueva string

	var cadenaTemp = strings.Fields(str)
	for pos,cad := range cadenaTemp {
		if( len(cad) > 4 ){
			var cadTemp = []rune(cadenaTemp[pos])
			for i, j := 0, len(cadTemp) - 1; i < j; i, j = i + 1, j - 1{
				cadTemp[i], cadTemp[j] = cadTemp[j], cadTemp[i]
			}
			cadenaTemp[pos] = string(cadTemp)
		}
	}

	cadenaNueva = strings.Join(cadenaTemp, " ")
	return cadenaNueva
}

func main(){
	fmt.Println( spinWords("hola mundo") )
}