// 2311102037 Brian Farrel Evandhika IF 11 06
package main

import (
	"fmt"
	"sort"
)

func median(arr []int) float64 { // fungsi untuk menghitung median
	n := len(arr)
	sort.Ints(arr) // fungsi untuk sorting

	if n%2 == 0 { // menghitung median untuk jumlah elemen genap
		return float64(arr[n/2-1]+arr[n/2]) / 2.0
	} else { // menghitung median untuk jumlah elemen ganjil
		return float64(arr[n/2])
	}
}

func main() { // fungsi utama untuk memanggil fungsi lainnya
	var input int
	var data []int

	fmt.Println("masukkan bilangan bulat positif satu per satu. Masukkan 0 untuk mencetak median, atau -5313541 untuk mengakhiri input.")

	for {
		fmt.Print("masukkan bilangan: ")
		fmt.Scan(&input)

		if input == -5313541 {
			break
		} else if input == 0 {
			if len(data) > 0 {
				fmt.Printf("median: %.2f\n", median(data))
			} else {
				fmt.Println("tidak ada data untuk menghitung median.")
			}
		} else {
			data = append(data, input)
		}
	}

	fmt.Println("input selesai.")
}
