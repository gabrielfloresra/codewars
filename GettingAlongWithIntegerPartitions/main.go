package main

import (
	"fmt"
	"sort"
)

func sumArrayInt(array []int) int {
	sum := 0

	for _, val := range array {
		sum += val
	}
	return sum
}

func productoArray(array []int) int {
	p := 1
	for _, val := range array {
		p *= val
	}
	return p
}
func getKeysMap(productos map[int]string) []int {
	keys := make([]int, 0, len(productos))
	for key := range productos {
		keys = append(keys, key)
	}
	return keys
}
func appendZeros(array []int, num int) []int {

	for sumArrayInt(array) != num {
		array = append(array, 1)
	}

	return array
}

func appendNewArray(newArray [][]int, array []int) [][]int {
	na := make([]int, 0)
	for _, val := range array {
		na = append(na, val)
	}
	newArray = append(newArray, na)
	return newArray
}

func getArrays(pos int, a []int) [][]int {
	var newElements [][]int

	if len(a) > (pos+1) && a[pos-1] >= (a[pos]+a[pos+1]) {
		newA := make([]int, len(a))
		copy(newA, a)
		newA[pos] = newA[pos] + newA[pos+1]
		newA = newA[:len(newA)-1]
		newElements = appendNewArray(newElements, newA)
		arraysLeft := getArrays(pos+1, newA)
		if len(arraysLeft) > 0 {
			for _, val := range arraysLeft {
				newElements = append(newElements, val)
			}
		}
		arraysJoin := getArrays(pos, newA)
		if len(arraysJoin) > 0 {
			for _, val := range arraysJoin {
				newElements = append(newElements, val)
			}
		}
	}
	return newElements
}
func part(num int) string {
	rangeValue := 0
	averageValue := 0.0
	medianValue := 0.0

	result := [][]int{[]int{num}}
	productosArray := make(map[int]int)
	keysProductosArrray := []int{num}
	productosArray[num] = 1
	averageValue += float64(num)
	numEvaluado := num - 1

	for numEvaluado != 0 {

		newArray := []int{numEvaluado}

		// agrega 1's al arreglo :v
		// asi es prro, de 1's >:v
		newArray = appendZeros(newArray, num)
		a := make([]int, len(newArray))
		copy(a, newArray)
		result = append(result, a)
		productoRes := productoArray(a)
		keysProductosArrray = append(keysProductosArrray, productoRes)
		productosArray[productoRes] = 1
		averageValue += float64(productoRes)

		newArrays := [][]int{}
		newArrays = getArrays(1, a)
		if len(newArrays) > 0 {
			for _, val := range newArrays {
				result = append(result, val)
				productoRes := productoArray(val)
				_, exist := productosArray[productoRes]
				if !exist {
					productosArray[productoRes] = 0
					keysProductosArrray = append(keysProductosArrray, productoRes)
					averageValue += float64(productoRes)
				}
				productosArray[productoRes]++
			}
		}
		numEvaluado--
	}

	sort.Ints(keysProductosArrray)
	averageValue /= float64(len(keysProductosArrray))
	rangeValue = keysProductosArrray[len(keysProductosArrray)-1] - keysProductosArrray[0]
	if len(keysProductosArrray)%2 == 0 {
		medianValue = float64(keysProductosArrray[len(keysProductosArrray)/2] + keysProductosArrray[(len(keysProductosArrray)/2)-1])
		medianValue /= 2
	} else {
		medianValue = float64(keysProductosArrray[len(keysProductosArrray)/2])
	}

	var resultValues string
	resultValues = fmt.Sprintf("Range: %d Average: %.2f Median: %.2f", rangeValue, averageValue, medianValue)
	return resultValues
}

func main() {
	//fmt.Println(enum(4))
	fmt.Println(part(30))
}
