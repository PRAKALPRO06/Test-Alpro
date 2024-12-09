package main

import (
	"fmt"
	"sort"
)

const NMAX = 1000000

// Struct untuk menyimpan nama partai dan jumlah suara
type Partai struct {
	nama  int
	suara int
}

// Fungsi untuk mencari indeks partai di dalam array
func posisi_2311102264(tabPartai []Partai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if tabPartai[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {
	var partai []Partai // Array untuk menyimpan data partai
	var n int           // Jumlah partai yang ditemukan

	// Membaca input pengguna
	fmt.Println("Masukkan data suara (akhiri dengan -1):")

	for {
		var suara int
		_, err := fmt.Scan(&suara)
		if err != nil {
			fmt.Println("Input tidak valid")
			return // Hentikan program jika terjadi error
		}
		if suara == -1 {
			break
		}
		idx := posisi_2311102264(partai, n, suara)
		if idx == -1 {
			// Jika partai belum ada di array, tambahkan partai baru
			partai = append(partai, Partai{nama: suara, suara: 1})
			n++
		} else {
			// Jika partai sudah ada, tambahkan jumlah suaranya
			partai[idx].suara++
		}
	}

	// Sort array berdasarkan suara (descending) dan nama (ascending jika suara sama)
	sort.Slice(partai, func(i, j int) bool {
		if partai[i].suara == partai[j].suara {
			return partai[i].nama < partai[j].nama
		}
		return partai[i].suara > partai[j].suara
	})

	// Output hasil
	fmt.Println("Hasil perolehan suara:")
	for _, p := range partai {
		fmt.Printf("%d(%d) ", p.nama, p.suara)
	}
	fmt.Println()
}
