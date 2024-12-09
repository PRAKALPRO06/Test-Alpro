package main

import (
	"fmt"
	"sort"
)

// Fungsi untuk menghitung median
func median(data_2311102024 []int) float64 {
	n := len(data_2311102024)
	if n == 0 {
		return 0
	}
	// Jika jumlah data ganjil, median adalah nilai tengah
	// Jika jumlah data genap, median adalah rata-rata dari dua nilai tengah
	if n%2 == 1 {
		return float64(data_2311102024[n/2])
	}
	return float64(data_2311102024[(n/2)-1]+data_2311102024[n/2]) / 2
}

func main() {
	const sentinel = -5313541 // Penanda akhir input
	var data_2311102024 []int
	var input int

	fmt.Println("Masukkan bilangan (akhiri dengan -5313541):")

	// Membaca data
	for {
		fmt.Scan(&input)
		if input == sentinel {
			break
		}
		if input == 0 {
			// Jika menemukan 0, hitung median
			sort.Ints(data_2311102024) // Mengurutkan data
			fmt.Printf("Median: %.2f\n", median(data_2311102024))
		} else {
			// Tambahkan bilangan ke dalam array
			data_2311102024 = append(data_2311102024, input)
		}
	}
}