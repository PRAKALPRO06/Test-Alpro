package main

import "fmt"

const MMAX = 1000000

type partai_2311102166 struct {
	nama string
	suara int
}

var tabPartai [MMAX]partai_2311102166
var n int // jumlah partai yang diinputkan

// Fungsi untuk mencari indeks suatu partai berdasarkan nama
func posisi(tabPartai []partai_2311102166, n int, nama string) int {
	for i := 0; i < n; i++ {
		if tabPartai[i].nama == nama {
			return i
		}
	}
	return -1
}

// Fungsi untuk mengurutkan data berdasarkan suara secara descending
func sortDescending(tabPartai []partai_2311102166, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if tabPartai[i].suara < tabPartai[j].suara {
				tabPartai[i], tabPartai[j] = tabPartai[j], tabPartai[i]
			}
		}
	}
}

func main() {
	var input string
	n = 0

	fmt.Println("Masukkan data suara partai (contoh: 1 1 2 3 3 3 4), akhiri dengan -1:")
	for {
		fmt.Scan(&input)
		if input == "-1" {
			break
		}
		idx := posisi(tabPartai[:], n, input)
		if idx != -1 {
			tabPartai[idx].suara++
		} else {
			tabPartai[n] = partai_2311102166{nama: input, suara: 1}
			n++
		}
	}

	// Mengurutkan partai berdasarkan suara
	sortDescending(tabPartai[:], n)

	// Menampilkan hasil
	fmt.Println("Hasil perolehan suara:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s(%d) ", tabPartai[i].nama, tabPartai[i].suara)
	}
	fmt.Println()
}
