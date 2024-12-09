package main

import (
	"fmt"
	"sort"
)

func main() {
	var numbers []int
	for {
		var input int
		fmt.Scan(&input)
		if input == -5313541 {
			break
		}
		if input != 0 {
			numbers = append(numbers, input)
		}
	}

	// mengurutkan nama
	sort.Ints(numbers)

	// menghitung median
	median := calculateMedian(numbers)

	// cetak
	fmt.Println("Median:", median)
}

func calculateMedian(data []int) float64 {
	length := len(data)
	if length == 0 {
		return 0
	}

	if length%2 == 1 {

		return float64(data[length/2])
	} else {

		mid1 := data[length/2-1]
		mid2 := data[length/2]
		return float64(mid1+mid2) / 2.0
	}
}
