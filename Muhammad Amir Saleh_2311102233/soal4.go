package main

import (
	"fmt"
)

const NMAX233 = 1000000

type arrInt [NMAX233]int

func sorting(T *arrInt, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if T[j] < T[minIdx] {
				minIdx = j
			}
		}
		T[i], T[minIdx] = T[minIdx], T[i]
	}
}

func median(T arrInt, n int) float64 {
	if n%2 == 0 {
		return float64(T[n/2-1]+T[n/2]) / 2.0
	}
	return float64(T[n/2])
}

func main() {
	var (
		data  arrInt
		count int
		input int
	)

	fmt.Println("Masukkan data (akhiri dengan -5313541):")
	for {
		fmt.Scan(&input)
		if input == -5313541 {
			break
		} else if input == 0 {
			sorting(&data, count)
			fmt.Printf("Median: %.2f\n", median(data, count))
		} else {
			data[count] = input
			count++
			if count >= NMAX233 {
				fmt.Println("Jumlah data melebihi kapasitas maksimum.")
				break
			}
		}
	}
}
