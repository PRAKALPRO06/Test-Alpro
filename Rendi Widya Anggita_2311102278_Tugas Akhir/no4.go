package main

import (
	"fmt"
	"sort"
)

func main() {
	var data_2311102278 []int // Array untuk menyimpan data input
	var input int

	fmt.Println("Masukkan angka (akhiri dengan -5313541):")

	for {
		// Membaca input angka
		fmt.Scan(&input)

		if input == -5313541 { // Marker akhir input
			break
		}

		if input != 0 { // Data selain 0 disimpan
			data_2311102278 = append(data_2311102278, input)
		} else {
			// Hitung median jika angka 0 ditemukan
			if len(data_2311102278) == 0 {
				fmt.Println("Data kosong. Tidak dapat menghitung median.")
				continue
			}

			// Urutkan data
			sort.Ints(data_2311102278)

			// Hitung median
			n := len(data_2311102278)
			median := 0
			if n%2 == 0 {
				median = (data_2311102278[n/2-1] + data_2311102278[n/2]) / 2
			} else {
				median = data_2311102278[n/2]
			}

			fmt.Println("Median saat ini:", median)
		}
	}

	fmt.Println("Program selesai.")
}
