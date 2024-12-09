package main

import (
	"fmt"
)

// Kapasitas maksimum array
const MAXBILL = 1000

// Array untuk menyimpan bilangan
type arrInt [MAXBILL]int

// Fungsi untuk mengurutkan array menggunakan selection sort
func sort_2311102241(T *arrInt, n int) {
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

// Fungsi untuk menghitung median array yang sudah diurutkan
func hitungmedian_2311100241(T arrInt, n int) float64 {
	if n%2 == 0 {
		// Jika jumlah elemen genap, median adalah rata-rata 2 elemen tengah
		return float64(T[n/2-1]+T[n/2]) / 2.0
	}
	// Jika jumlah elemen ganjil, median adalah elemen tengah
	return float64(T[n/2])
}

func main() {
	var T arrInt // Array untuk menyimpan bilangan
	var n int    // Jumlah elemen dalam array

	fmt.Println("Masukkan bilangan satu per satu (akhiri dengan -1):")

	for {
		var num int
		fmt.Print("Bilangan: ")
		fmt.Scan(&num)

		if num == -1 {
			// Jika input adalah -1, berhenti membaca
			break
		}

		// Tambahkan angka ke array
		T[n] = num
		n++
	}

	// Mengurutkan array
	sort_2311102241(&T, n)

	// Menghitung median
	median := hitungmedian_2311100241(T, n)

	// Menampilkan hasil
	fmt.Println("Array setelah diurutkan:", T[:n])
	fmt.Printf("Median: %.1f\n", median)
}
