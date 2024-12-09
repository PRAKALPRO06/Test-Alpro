package main

import (
	"fmt"
	"sort"
)

func main() {//di dalam fungsi ini main ini terdapat perulangan untuk meghitung median dan menambahaan angka ke array
	const END_MARKER = -5313541
	var numbers []int
	var input int

	for {
		fmt.Scan(&input)
		if input == END_MARKER {
			break
		}
		if input == 0 {
			if len(numbers) == 0 {
				fmt.Println(0)
			} else {
				median := menghitungmedian_2311102263(numbers)
				fmt.Printf("%.1f\n", median)
			}
		} else {
			numbers = append(numbers, input)
		}
	}
}

func menghitungmedian_2311102263(arr []int) float64 {//di dalam fungsi menghitung midean ini terdapat perintah untuk mengurutkan array dan duplikat dan menentukan jumlah data apakah ganjl maupun genap
	// Salin array dan urutkan
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)

	n := len(sortedArr)
	if n%2 == 1 {
		return float64(sortedArr[n/2])
	}
	return float64(sortedArr[n/2-1]+sortedArr[n/2]) / 2.0
}
