package main

import (
	"fmt"
)

// Definisi tipe khusus untuk daftar bilangan
type IntArray []int

// Method untuk menyisipkan angka_2311102273 ke dalam array (insertion sort)
func (arr *IntArray) insertAndSort(num int) {
	*arr = append(*arr, num) // Tambahkan angka_2311102273 ke array
	for i := 1; i < len(*arr); i++ {
		key := (*arr)[i]
		j := i - 1

		// Geser elemen ke kanan untuk menjaga urutan
		for j >= 0 && (*arr)[j] > key {
			(*arr)[j+1] = (*arr)[j]
			j--
		}
		(*arr)[j+1] = key
	}
}

// Method untuk menghitung median dari IntArray
func (arr IntArray) median() int {
	n := len(arr)
	if n%2 == 1 {
		// Jika jumlah elemen ganjil
		return arr[n/2]
	}
	// Jika jumlah elemen genap
	return (arr[n/2-1] + arr[n/2]) / 2
}

func main() {
	var angka_2311102273 IntArray
	var input int

	fmt.Println("Masukkan bilangan satu per satu (0 untuk menghitung median, -5313 untuk keluar):")

	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error membaca input:", err)
			break
		}

		if input == -5313 {
			// Keluar dari program
			break
		}

		if input == 0 {
			// Cetak median jika angka_2311102273 0
			if len(angka_2311102273) > 0 {
				fmt.Println("Median saat ini:", angka_2311102273.median())
			}
			continue
		}

		// Tambahkan angka_2311102273 ke array dan sortir
		angka_2311102273.insertAndSort(input)
	}
}
