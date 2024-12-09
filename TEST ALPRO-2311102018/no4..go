package main

import (
	"fmt"
	"sort"
)

// Fungsi menghitung median
func median_2311102018(data []int) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}
	// Apabila seluruh data ganjil, median adalah nilai tengah
	// Jika seluruh data genap, median adalah rata-rata dari dua nilai tengah
	if n%2 == 1 {
		return float64(data[n/2])
	}
	return float64(data[(n/2)-1]+data[n/2]) / 2
}

func main() {
	const sentinel = -5313541 //  akhir input
	var data []int
	var input int

	fmt.Println("Masukkan bilangan (akhiri dengan -5313541):")

	//data terbaca apabila data input sama dengan akhir input maka akan berhenti/brake
	for {
		fmt.Scan(&input)
		if input == sentinel {
			break
		}
		if input == 0 {
			// Jika menemukan 0, hitung median
			sort.Ints(data) // Mengurutkan data
			fmt.Printf("Median: %.2f\n", median_2311102018(data))
		} else {
			// Tambahkan bilangan ke dalam array
			data = append(data, input)
		}
	}
}