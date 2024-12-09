package main

import (
	"fmt"
)

const NMAX = 1000000

// Struct untuk menyimpan nama partai dan jumlah suara
type partai struct {
	nama  int
	suara int
}

// Array of partai dengan kapasitas maksimum
type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, nama int) int {
	// Mencari posisi partai berdasarkan nama dengan pencarian sekuensial
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1 // Mengembalikan -1 jika tidak ditemukan
}

func insertionSortDescending(t *tabPartai, n int) {
	// Proses pengurutan dengan insertion sort secara descending berdasarkan jumlah suara
	for i := 1; i < n; i++ {
		key := t[i]
		j := i - 1
		for j >= 0 && t[j].suara < key.suara {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = key
	}
}

func main() {
	var p_2311102273 tabPartai
	var n int // Jumlah partai yang terdata

	fmt.Println("Masukkan suara partai (akhiri dengan -1):")
	for {
		var nama int
		fmt.Scan(&nama)
		if nama == -1 {
			break
		}

		pos := posisi(p_2311102273, n, nama)
		if pos == -1 {
			// Tambahkan partai baru jika belum ada
			p_2311102273[n] = partai{nama: nama, suara: 1}
			n++
		} else {
			// Tambahkan suara ke partai yang sudah ada
			p_2311102273[pos].suara++
		}
	}

	// Urutkan array p_2311102273 berdasarkan suara secara descending
	insertionSortDescending(&p_2311102273, n)

	// Tampilkan hasil
	fmt.Println("Hasil perolehan suara:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d)\n", p_2311102273[i].nama, p_2311102273[i].suara)
	}
}
