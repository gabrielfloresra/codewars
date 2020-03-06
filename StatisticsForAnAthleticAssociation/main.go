package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func toSeg(h int, m int, s int) int {
	return (h * 3600) + (m * 60) + s
}

func formatSeg(seg int) (int, int, int) {
	s := seg % 60
	seg /= 60
	m := seg % 60
	h := seg / 60

	return h, m, s
}

func getRangeValues(a []int, b []int) (int, int, int) {
	segA := toSeg(a[0], a[1], a[2])
	segB := toSeg(b[0], b[1], b[2])

	temp := segB - segA

	return formatSeg(temp)
}

func getMedianValues(mapTimes [][]int) (int, int, int) {
	pos := len(mapTimes) / 2
	s := toSeg(mapTimes[pos][0], mapTimes[pos][1], mapTimes[pos][2])
	if len(mapTimes)%2 == 0 {
		s += toSeg(mapTimes[pos+1][0], mapTimes[pos+1][1], mapTimes[pos+1][2])
		s /= 2
	}
	return formatSeg(s)
}

func stati(cad string) string {

	cad = strings.Replace(cad, " ", "", -1)
	if len(cad) == 0 {
		return ""
	}
	times := strings.Split(cad, ",")

	var mapTimes [][]int
	segAverage := 0

	for pos, val := range times {
		data := strings.Split(val, "|")

		var h, m, s int
		var err error
		h, err = strconv.Atoi(data[0])
		if m, err = strconv.Atoi(data[1]); err != nil {
			m = 0
		}
		if s, err = strconv.Atoi(data[2]); err != nil {
			s = 0
		}

		segAverage += s + (m * 60) + (h * 3600)

		mapTimes = append(mapTimes, make([]int, 3))
		mapTimes[pos][0] = h
		mapTimes[pos][1] = m
		mapTimes[pos][2] = s
	}

	sort.SliceStable(mapTimes, func(i, j int) bool {
		if mapTimes[i][0] == mapTimes[j][0] {
			if mapTimes[i][1] == mapTimes[j][1] {
				return mapTimes[i][2] < mapTimes[j][2]
			}
			return mapTimes[i][1] < mapTimes[j][1]
		}
		return mapTimes[i][0] < mapTimes[j][0]
	})

	// operaciones para obtener el rango
	hRange, mRange, sRange := getRangeValues(mapTimes[0], mapTimes[len(mapTimes)-1])
	rangeValues := []int{hRange, mRange, sRange}

	// operaciones para obtener la mediana
	hMedian, mMedian, sMedian := getMedianValues(mapTimes)
	medianValues := []int{hMedian, mMedian, sMedian}

	// operaciones para obtener el promedio
	segAverage = segAverage / len(mapTimes)
	hAverage, mAverage, sAverage := formatSeg(segAverage)
	averageValues := []int{hAverage, mAverage, sAverage}

	var cadFinal string
	cadFinal = fmt.Sprintf("Range: %02d|%02d|%02d Average: %02d|%02d|%02d Median: %02d|%02d|%02d", rangeValues[0], rangeValues[1], rangeValues[2], averageValues[0], averageValues[1], averageValues[2], medianValues[0], medianValues[1], medianValues[2])

	return cadFinal
}

func main() {
	fmt.Println(stati("02|15|30, 1|45|30"))
}
