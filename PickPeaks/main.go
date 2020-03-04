package main

import "fmt"

//PosPeaks sa
type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func peaks(array []int) PosPeaks {

	// var pp PosPeaks
	// pp.Pos = []int{}
	// pp.Peaks = []int{}
	// var posPrev = 0

	// if len(array) == 0 {
	// 	return pp
	// }

	// for posPrev < len(array)-1 {
	// 	var posNext = posPrev + 1
	// 	var posInit = -1
	// 	for posPrev < len(array)-1 && array[posPrev] <= array[posNext] {
	// 		if array[posPrev] == array[posNext] && posInit == -1 {
	// 			posInit = posPrev
	// 		}
	// 		posPrev, posNext = posNext, posNext+1
	// 	}
	// 	if array[posInit] != array[posPrev] {
	// 		posInit = -1
	// 	}
	// 	if posPrev < len(array)-1 {
	// 		if posPrev != 0 && posInit != 0 {
	// 			if posInit != -1 {
	// 				pp.Pos = append(pp.Pos, posInit)
	// 			} else {
	// 				pp.Pos = append(pp.Pos, posPrev)
	// 			}
	// 			pp.Peaks = append(pp.Peaks, array[posPrev])
	// 		}
	// 		for posPrev < len(array)-1 && array[posPrev] >= array[posNext] {
	// 			posPrev++
	// 			posNext++
	// 		}
	// 	}
	// }
	// return pp
	posPeaks := PosPeaks{Pos: []int{}, Peaks: []int{}}
	candidate := 0
	for i := 1; i < len(array); i++ {
		if array[i-1] < array[i] {
			candidate = i
		} else if array[i-1] > array[i] && candidate > 0 {
			posPeaks.Pos = append(posPeaks.Pos, candidate)
			posPeaks.Peaks = append(posPeaks.Peaks, array[candidate])
			candidate = 0
		}
	}
	return posPeaks
}

func main() {
	// fmt.Print(peaks([]int{13, 9, -2, -5, 8, 8, 14, -2, -3}))
	fmt.Print(peaks([]int{1, 2, 2, 2, 1}))
}
