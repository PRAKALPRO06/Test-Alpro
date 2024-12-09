package main

import "fmt"

// Fungsi untuk mengecek apakah elemen sudah ada dalam slice
func exist(T []int, val int) bool {
	for _, v := range T {
		if v == val {
			return true
		}
	}
	return false
}

// Fungsi untuk mengisi slice dengan bilangan unik hingga duplikat ditemukan
func inputSet() []int {
	var T []int
	var val int
	for {
		fmt.Scan(&val)
		if exist(T, val) { // Hentikan jika bilangan sudah ada
			break
		}
		T = append(T, val)
	}
	return T
}

// Fungsi untuk mencari irisan dari dua slice
func findIntersection(T1, T2 []int) []int {
	var T3 []int
	for _, v := range T1 {
		if exist(T2, v) && !exist(T3, v) { // Tambahkan ke irisan jika belum ada
			T3 = append(T3, v)
		}
	}
	return T3
}

// Fungsi untuk mencetak slice secara horizontal
func printSet(T []int) {
	for i, v := range T {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func main() {
	fmt.Println("Masukkan himpunan pertama (berhenti jika duplikat):")
	s1 := inputSet()

	fmt.Println("Masukkan himpunan kedua (berhenti jika duplikat):")
	s2 := inputSet()

	// Cari irisan kedua himpunan
	s3 := findIntersection(s1, s2)

	// Cetak hasil irisan
	fmt.Println("Hasil irisan:")
	printSet(s3)
}
