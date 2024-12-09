package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Program ini dibuat oleh [NIM 2311102276]
func main() {
	fmt.Println("Program Median - Dibuat oleh NIM 2311102276")

	scanner := bufio.NewScanner(os.Stdin)
	var numbers []int

	fmt.Println("Masukkan angka (masukkan -5313541 untuk berhenti):")
	for scanner.Scan() {
		input := scanner.Text()

		// Konversi input ke integer
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Masukkan angka yang valid!")
			continue
		}

		// Jika input adalah -5313541, hentikan masukan
		if number == -5313541 {
			break
		}

		// Tambahkan angka ke daftar
		numbers = append(numbers, number)

		// Hitung median
		median := calculateMedian(numbers)
		fmt.Printf("Median saat ini: %.2f\n", median)
	}
}

// Fungsi untuk menghitung median
func calculateMedian(numbers []int) float64 {
	sort.Ints(numbers) // Urutkan data

	n := len(numbers)
	if n%2 == 1 {
		// Jika jumlah data ganjil, median adalah elemen tengah
		return float64(numbers[n/2])
	}
	// Jika jumlah data genap, median adalah rata-rata dua elemen tengah
	return float64(numbers[n/2-1]+numbers[n/2]) / 2
}
