package main

import (
	"fmt"
)

type set [2024]int

// Memeriksa apakah suatu nilai (val) sudah ada dalam himpunan `T` hingga indeks tertentu (n).
func exist(T set, n int, val233 int) bool {
	for i := 0; i < n; i++ {
		if T[i] == val233 {
			return true
		}
	}
	return false
}

// Membaca input dari pengguna dan mengisi himpunan `T` dengan bilangan bulat hingga ditemukan bilangan bulat yang sama.
func inputSet(T *set, n *int) {
	var val233 int
	*n = 0
	for {
		fmt.Scan(&val233)
		if exist(*T, *n, val233) {
			break
		}
		(*T)[*n] = val233
		*n++
	}
}

// Menemukan irisan dari dua himpunan (`T1` dan `T2`) dan menyimpannya ke dalam himpunan hasil `T3`.
func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	*h = 0
	for i := 0; i < n; i++ {
		if exist(T2, m, T1[i]) {
			(*T3)[*h] = T1[i]
			*h++
		}
	}
}

// Mencetak bilangan bulat dari himpunan `T` hingga indeks tertentu (n).
func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int
	fmt.Println("Masukkan himpunan pertama (akhiri dengan duplikat):")
	inputSet(&s1, &n1)
	fmt.Println("Masukkan himpunan kedua (akhiri dengan duplikat):")
	inputSet(&s2, &n2)
	findIntersection(s1, s2, n1, n2, &s3, &n3)
	fmt.Println("Irisan kedua himpunan adalah:")
	printSet(s3, n3)
}
