package main

import (
	"fmt"
)

type set [2022]int

// Memeriksa apakah nilai sudah ada di array
func exist(T set, n int, val int) bool {
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

// Memasukkan nilai ke dalam array himpunan
func inputSet(T *set, n *int) {
	var val int
	*n = 0 // Inisialisasi jumlah elemen
	fmt.Println("Masukkan elemen himpunan (berhenti jika duplikat ditemukan):")
	for {
		fmt.Scan(&val)
		if exist(*T, *n, val) { // Jika duplikat ditemukan, berhenti
			fmt.Println("Duplikat ditemukan, input dihentikan.")
			break
		}
		T[*n] = val
		*n++
	}
}

// Mencari irisan dua himpunan
func findIntersection(T1, T2 set, n1, n2 int, T3 *set, n3 *int) {
	*n3 = 0 // Inisialisasi jumlah elemen irisan
	for i := 0; i < n1; i++ {
		if exist(T2, n2, T1[i]) && !exist(*T3, *n3, T1[i]) {
			T3[*n3] = T1[i]
			(*n3)++
		}
	}
}

// Menampilkan elemen-elemen dalam himpunan
func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()
}

func main() {
	var s1, s2, s3_2311102276 set // Variabel s3 diberi nama sesuai NIM
	var n1, n2, n3 int

	// Input himpunan pertama
	fmt.Println("Input himpunan pertama:")
	inputSet(&s1, &n1)

	// Input himpunan kedua
	fmt.Println("Input himpunan kedua:")
	inputSet(&s2, &n2)

	// Cari irisan
	findIntersection(s1, s2, n1, n2, &s3_2311102276, &n3)

	// Cetak hasil irisan
	fmt.Println("Irisan himpunan:")
	printSet(s3_2311102276, n3)
}
