package main

import (
	"fmt"
	"sort"
)

// Fungsi untuk menghitung median
func median(Data_2311102033 []int) float64 {
	n := len(Data_2311102033)
	if n == 0 {
		return 0
	}
	// Jika jumlah Data ganjil, median adalah nilai tengah
	// Jika jumlah Data genap, median adalah rata-rata dari dua nilai tengah
	if n%2 == 1 {
		return float64(Data_2311102033[n/2])
	}
	return float64(Data_2311102033[(n/2)-1]+Data_2311102033[n/2]) / 2
}

func main() {
	const sentinel = -5313541 // Penanda akhir input
	var Data_2311102033 []int
	var input int

	fmt.Println("Masukkan bilangan (akhiri dengan -5313541):")

	// Membaca Data
	for {
		fmt.Scan(&input)
		if input == sentinel {
			break
		}
		if input == 0 {
			// Jika menemukan 0, hitung median
			sort.Ints(Data_2311102033) // Mengurutkan
			fmt.Printf("Median: %.2f\n", median(Data_2311102033))
		} else {
			// Tambahkan bilangan ke dalam array
			Data_2311102033 = append(Data_2311102033, input)
		}
	}
}
