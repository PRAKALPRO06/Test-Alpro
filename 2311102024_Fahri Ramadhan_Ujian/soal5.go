package main

import
	"fmt"

const NMAX = 1000000

// Struktur data partai
type Partai struct {
	nama  int
	suara int
}

// Tipe array untuk menyimpan data partai
type TabPartai_2311102024 [NMAX]Partai

func main() {
	var t TabPartai_2311102024
	var n int // Jumlah partai yang diproses
	var input int

	// Inisialisasi jumlah partai
	n = 0

	// Input suara partai
	for {
		fmt.Scan(&input)
		if input == -1 {
			break
		}

		// Cari posisi partai berdasarkan nama
		pos := posisi(t, n, input)

		if pos == -1 {
			// Jika partai belum ada, tambahkan partai baru
			t[n].nama = input
			t[n].suara = 1
			n++
		} else {
			// Jika partai sudah ada, tambahkan jumlah suaranya
			t[pos].suara++
		}
	}

	// Pengurutan secara descending berdasarkan jumlah suara
	insertionSortDescending(&t, n)

	// Tampilkan hasil
	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d) ", t[i].nama, t[i].suara)
	}
	fmt.Println()
}

// Fungsi untuk mencari indeks partai berdasarkan nama
func posisi(t TabPartai_2311102024, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

// Fungsi untuk mengurutkan array partai secara descending berdasarkan jumlah suara
func insertionSortDescending(t *TabPartai_2311102024, n int) {
	for i := 1; i < n; i++ {
		key := t[i]
		j := i - 1

		// Pindahkan elemen yang lebih kecil ke kanan
		for j >= 0 && t[j].suara < key.suara {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = key
	}
}