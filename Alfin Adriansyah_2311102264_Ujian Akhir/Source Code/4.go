package main

import (
	"fmt"
)

const NMAX = 1000000 // Batas maksimum array

type arrInt [NMAX]int

// Fungsi untuk mengurutkan array menggunakan algoritma selection sort
func sorting_2311102260(T *arrInt, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[minIdx] {
				minIdx = j
			}
		}
		// Menukar elemen
		T[i], T[minIdx] = T[minIdx], T[i]
	}
}

// Fungsi untuk menghitung median dari array yang sudah diurutkan
func median_2311102260(T arrInt, n int) float64 {
	if n%2 == 0 {
		return float64(T[n/2-1]+T[n/2]) / 2.0
	}
	return float64(T[n/2])
}

func main() {
	var T arrInt
	n := 0
	step := 1

	fmt.Println("Masukkan data (pisahkan dengan spasi, akhiri dengan -5313541):")

	for {
		var num int
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("Error membaca input:", err)
			return // Hentikan program jika terjadi error
		}

		if num == 0 {
			// Jika input adalah 0, cetak hasil saat ini
			fmt.Printf("Sampai bilangan 0 yang ke-%d, data terbaca adalah ", step)
			step++
			fmt.Print(T[:n], ", setelah tersusun: ")
			sorting_2311102260(&T, n)
			fmt.Print(T[:n])
			fmt.Printf(", maka median saat itu adalah %.1f\n", median_2311102260(T, n))
		} else if num == -5313541 {
			break // Penanda akhir input
		} else {
			// Tambahkan angka ke array
			T[n] = num
			n++
		}
	}
}
