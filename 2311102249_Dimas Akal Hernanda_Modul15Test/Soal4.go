package main

import (
	"fmt"
)

const NMAX = 1000000

type arrInt [NMAX]int

// Fungsi untuk melakukan selection sort
func sorting(T *arrInt, n int) {
	/* I.S. T mendefinisikan n adalah bilangan bulat
	   F.S. Array T berurut membesar menggunakan algoritma selection sort */
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[minIdx] {
				minIdx = j
			}
		}
		T[i], T[minIdx] = T[minIdx], T[i]
	}
}

// Fungsi untuk menghitung median dari array yang sudah terurut
func median(T arrInt, n int) float64 {
	if n%2 == 1 {

		return float64(T[n/2])
	} else {

		return float64(T[n/2-1]+T[n/2]) / 2
	}
}

func main() {
	var data arrInt
	var n int
	var x int

	fmt.Println("Masukkan bilangan (akhiri dengan -5313541)_2311102249:")
	for {
		fmt.Scan(&x)

		// Jika marker -5313541 ditemukan, hentikan program
		if x == -5313541 {
			break
		}

		// Jika data adalah 0, urutkan dan hitung median
		if x == 0 {
			// Mengurutkan data menggunakan selection sort
			sorting(&data, n)

			// Menghitung median
			med := median(data, n)
			fmt.Printf("%.1f\n", med) // Cetak median
		} else {
			// Menyimpan bilangan positif ke dalam array
			data[n] = x
			n++
		}
	}
}
