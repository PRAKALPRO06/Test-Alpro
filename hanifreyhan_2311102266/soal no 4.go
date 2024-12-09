package main

import (
	"fmt"
	"sort"
)

// Fungsi untuk menghitung median
func Media2311102266(numbers []int) float64 {
	sort.Ints(numbers)
	n := len(numbers)

	if n%2 == 0 {
		mid1 := numbers[n/2-1]
		mid2 := numbers[n/2]
		return float64(mid1+mid2) / 2.0
	} else {
		return float64(numbers[n/2])
	}
}

func main() {
	// masukan data
	input := []int{7, 23, 11, 0, 5, 19, 2, 29, 3, 13, 17, 0, -5313541}
	var data []int



	for _, num := range input {
		if num == 0 {
			// Hitung median saat menemukan angka 0
			if len(data) > 0 {
				median := Media2311102266(data)
				fmt.Printf("%.1f\n", median)
			}
		} else if num == -5313541 {
			// Akhiri saat menemukan marker
			break
		} else {
			// Tambahkan angka ke data
			data = append(data, num)
		}
	}
}
