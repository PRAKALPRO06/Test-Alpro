package main

import (
	"fmt"
	"sort"
)

// Fungsi untuk menghitung median, ketika jumlah datanya ganjil maka median adalah nilai tengahnya, namun ketika jumlah genap median adalah rata rata dari nilai tengah
func median(data []int) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}
	if n%2 == 1 {
		return float64(data[n/2])
	}
	return float64(data[(n/2)-1]+data[n/2]) / 2
}

func main() {
	const sentinel = -5313541 
	var data_2311102011 []int
	var input int

	fmt.Println("Masukkan bilangan (akhiri dengan -5313541):")

	for {
		fmt.Scan(&input)
		if input == sentinel {
			break
		}
		if input == 0 {
			sort.Ints(data_2311102011)
			fmt.Printf("Median: %.2f\n", median(data_2311102011))
		} else {
			data_2311102011 = append(data_2311102011, input)
		}
	}
}