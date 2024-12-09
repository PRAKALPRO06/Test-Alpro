package main

import (
	"fmt"
)

func main() {
	// Input dua baris himpunan
	var set1, set2 []int

	fmt.Println("Masukkan elemen baris pertama (akhiri dengan angka duplikat):")
	set1 = readSet()

	fmt.Println("Masukkan elemen baris kedua (akhiri dengan angka duplikat):")
	set2 = readSet()

	// Temukan irisan dari kedua himpunan
	intersection := findIntersection(set1, set2)

	// Cetak hasil
	fmt.Println("Irisan dari kedua himpunan:", intersection)
}

func readSet() []int {
	var input int
	set := make(map[int]bool) // Untuk memastikan tidak ada duplikat
	var result []int

	for {
		fmt.Scan(&input)
		if set[input] {
			break // Berhenti jika elemen sudah ada dalam himpunan
		}
		set[input] = true
		result = append(result, input)
	}

	return result
}

func findIntersection(set1, set2 []int) []int {
	setMap := make(map[int]bool)
	for _, num := range set1 {
		setMap[num] = true
	}

	var intersection []int
	for _, num := range set2 {
		if setMap[num] {
			intersection = append(intersection, num)
		}
	}

	return intersection
}