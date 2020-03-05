package main

import (
	"fmt"
	"math"
	"sort"
)

func comp(array1 []int, array2 []int) bool {

	if array1 == nil && array2 == nil {
		return false
	}
	if len(array1) == 0 && len(array2) == 0 && len(array1) == 0 {
		return true
	}
	if array1 == nil || len(array1) == 0 || array2 == nil || len(array2) == 0 || len(array1) != len(array2) {
		return false
	}

	sort.Ints(array1)
	sort.Ints(array2)

	var numMap map[int64]int
	numMap = make(map[int64]int)

	for _, val := range array1 {
		val := math.Abs(float64(val))
		_, exist := numMap[int64(val)]
		if !exist {
			numMap[int64(val)] = 1
		} else {
			numMap[int64(val)]++
		}
	}

	for _, val := range array2 {
		if val < 0 {
			return false
		}
		pos := int64(math.Sqrt(float64(val)))
		if math.Pow(float64(pos), float64(2)) != float64(val) {
			return false
		}
		num, exists := numMap[pos]
		if !exists || num == 0 {

			return false
		}
		numMap[pos]--
	}
	return true
}

func main() {

	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var a2 = []int{231, 14641, 20736, 361, 25921, 361, 20736, 361}
	fmt.Println(comp(a1, a2))

	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(comp(a1, a2))
	//dotest(a1, a2, true)
	a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	a2 = []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(comp(a1, a2))
	//dotest(a1, a2, false)
	a1 = nil
	a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(comp(a1, a2))
	//dotest(a1, a2, false)
}
