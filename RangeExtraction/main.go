package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Solution Function
func Solution(list []int) (result string) {

	for i, j := 0, 0; i < len(list); i = j {
		result += strconv.Itoa(list[i])
		for j = i; j < len(list)-1 && list[j+1]-list[j] == 1; j++ {
		}
		if j-i > 1 {
			result += "-"
		} else {
			result += ","
			if j == i {
				j++
			}
		}
	}
	return strings.Trim(result, ",")

	//	for i := 0; i < len(list); i++ {
	//		menor := list[i]
	//		num := menor
	//		j := i + 1
	//		for j < len(list) && list[j]-num == 1 {
	//			num = list[j]
	//			j++
	//		}
	//		condicion := num - menor
	//		result += "," + strconv.Itoa(menor)
	//		if condicion > 0 {
	//			if condicion == 1 {
	//				result += ","
	//			} else {
	//				result += "-"
	//			}
	//			result += strconv.Itoa(num)
	//		}
	//		i = j - 1
	//	}
	//
	//return strings.Trim(result, ",")
}
func main() {
	fmt.Println(Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
}
