// Geranada Saputra Priambudi
// 2311102008
// IF-11-06

package main

import "fmt"

// NMAX = batas maksimum jumlah data yang dapat disimpan
const NMAX = 1000000

// arrInt = tipe array yang berisi NMAX bilangan bulat
type arrInt [NMAX]int

// Fungsi sorting
// I.S. T adalah array dengan n elemen yang belum terurut
// F.S. T menjadi array dengan elemen-elemen terurut secara menaik menggunakan algoritma selection sort
func sorting(T *arrInt, n int) {
	var pass_2311102008, idx_min, i, temp int
	// Iterasi untuk setiap elemen array
	for pass_2311102008 = 1; pass_2311102008 <= n-1; pass_2311102008++ {
		idx_min = pass_2311102008 - 1
		for i = pass_2311102008; i <= n-1; i++ {
			if T[idx_min] > T[i] {
				idx_min = i
			}
		}
		temp = T[idx_min]
		T[idx_min] = T[pass_2311102008-1]
		T[pass_2311102008-1] = temp
	}
}

// Fungsi median
// Mengembalikan median dari array T yang sudah terurut berisi n elemen
func median(T arrInt, n int) float64 {
	var mid int = n / 2 // Indeks tengah array
	if n%2 == 0 {
		// Jika jumlah elemen genap, median adalah rata-rata dua elemen tengah
		return float64(T[mid-1]+T[mid]) / 2.0
	} else {
		// Jika jumlah elemen ganjil, median adalah elemen tengah
		return float64(T[mid])
	}
}

// Fungsi utama (main)
func main() {
	var A arrInt
	var x, n int
	n = 0
	fmt.Scan(&x)

	// Proses setiap bilangan sampai menemukan marker -5313541 atau mencapai batas NMAX
	for x != -5313541 && n < NMAX {
		if x == 0 {
			sorting(&A, n)
			fmt.Println(median(A, n))
		} else {
			// Tambahkan bilangan ke array A
			A[n] = x
			n++
		}
		// Baca bilangan berikutnya
		fmt.Scan(&x)
	}
}
