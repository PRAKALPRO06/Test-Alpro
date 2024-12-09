package main

import (
	"fmt"
	"sort"
)

// Fungsi untuk menghitung nilai median dari sebuah data 
func median(data []int) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}else if n%2 == 1 {
		return float64(data[n/2])
	}
	return float64(data[(n/2)-1]+data[n/2]) / 2
}

func main() {
	// membuat tanda untuk mengahiri inputan
	const sentinel = -5313541 
	var data []int
	var input int

	fmt.Println("Masukkan bilangan (akhiri dengan -5313541):")

	// Membaca data array
	for {
		fmt.Scan(&input)
		if input == sentinel {
			break
		}
		if input == 0 {
			sort.Ints(data) 
			fmt.Printf("Median: %.2f\n", median(data))
		} else {
			data = append(data, input)
		}
	}
}