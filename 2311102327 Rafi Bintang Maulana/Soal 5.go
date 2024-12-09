package main

import (
	"fmt"
)

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

//Fungsi cari posisi dengan perulangan
func posisi_2311102327(tabPartai []partai, n, nama int) int {
	for i := 0; i < n; i++ {
		if tabPartai[i].nama == nama {
			return i
		}
	}
	return -1
}

// Fungsi pengurutan
func insertionSort(tabPartai []partai, n int) {
	for i := 1; i < n; i++ {
		key := tabPartai[i]
		j := i - 1
		for j >= 0 && tabPartai[j].suara < key.suara {
			tabPartai[j+1] = tabPartai[j]
			j--
		}
		tabPartai[j+1] = key
	}
}

func main() {
	var tabPartai []partai
	var input, n int
	n = 0

	fmt.Println("Masukkan data suara (akhiri dengan -1 kalau mau end):")
	for {
		fmt.Scan(&input)
		if input == -1 {
			break
		}
		pos := posisi_2311102327(tabPartai, n, input)
		if pos == -1 {
			tabPartai = append(tabPartai, partai{nama: input, suara: 1})
			n++
		} else {
			tabPartai[pos].suara++
		}
	}

	insertionSort(tabPartai, n)

	fmt.Println("\nHasil perolehan suara:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d) ", tabPartai[i].nama, tabPartai[i].suara)
	}
	fmt.Println()
}
