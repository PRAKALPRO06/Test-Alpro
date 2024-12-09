package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []int

	fmt.Println("Masukkan angka satu per satu. Ketik '-5313541' untuk berhenti:")

	for {
		fmt.Print("Masukkan angka: ")
		if scanner.Scan() {
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
		} else {
			fmt.Println("Terjadi kesalahan input. Silakan coba lagi.")
			break
		}
	}
	fmt.Println("Program selesai. Terima kasih!")
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