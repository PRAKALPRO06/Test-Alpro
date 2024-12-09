package main

import (
	"fmt"
)

type set [2022]int

// Fungsi ada untuk mengecek apakah elemen val ada di dalam himpunan T.
func ada(T set, n int, val int) bool {
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

// Fungsi inputSet untuk menerima masukan hingga ada elemen duplikat atau himpunan penuh.
func inputSet(T *set, n *int) {
	var val int
	*n = 0
	seen := make(map[int]bool) // Untuk mengecek duplikat dengan cepat

	for {
		fmt.Scan(&val)
		if seen[val] {
			break
		}
		if *n >= len(T) {
			break
		}
		T[*n] = val
		*n++
		seen[val] = true
	}
}

// Fungsi temukanIrisan untuk mencari irisan antara dua himpunan T1 dan T2.
func temukanIrisan(T1, T2 set, n1, n2 int, T3 *set, n3 *int) {
	*n3 = 0
	for i := 0; i < n1; i++ {
		if ada(T2, n2, T1[i]) && !ada(*T3, *n3, T1[i]) {
			T3[*n3] = T1[i]
			(*n3)++
		}
	}
}

// Fungsi printSet untuk mencetak elemen dalam himpunan.
func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()
}

func main() {
	var s2311102273, s2, s3 set
	var n1, n2, n3 int

	// Memasukkan elemen untuk himpunan pertama
	fmt.Println("Masukkan elemen untuk himpunan pertama:")
	inputSet(&s2311102273, &n1)

	// Memasukkan elemen untuk himpunan kedua
	fmt.Println("Masukkan elemen untuk himpunan kedua :")
	inputSet(&s2, &n2)

	// Menemukan irisan dari himpunan pertama dan kedua
	temukanIrisan(s2311102273, s2, n1, n2, &s3, &n3)

	// Mencetak irisan himpunan
	fmt.Println("Irisan dari kedua himpunan adalah:")
	printSet(s3, n3)
}
