package main

import (
	"fmt"
	"sort"
)

// Struct untuk menyimpan data partai
type Partai_2311102033 struct {
	Nama  int
	Suara int
}

// Fungsi utama
func main() {
	var (
		input     int
		partaiMap = make(map[int]int) // Map untuk menyimpan jumlah suara per partai
	)

	fmt.Println("Masukkan angka partai yang dipilih (akhiri dengan -1):")

	// Membaca input hingga -1
	for {
		fmt.Scan(&input)
		if input == -1 {
			break
		}
		partaiMap[input]++
	}

	// Menyimpan data ke dalam slice untuk diurutkan
	var partaiList []Partai_2311102033
	for nama, suara := range partaiMap {
		partaiList = append(partaiList, Partai_2311102033{Nama: nama, Suara: suara})
	}

	// Mengurutkan slice berdasarkan jumlah suara secara descending
	sort.Slice(partaiList, func(i, j int) bool {
		if partaiList[i].Suara == partaiList[j].Suara {
			return partaiList[i].Nama < partaiList[j].Nama // Jika suara sama, urutkan berdasarkan nama partai
		}
		return partaiList[i].Suara > partaiList[j].Suara
	})

	// Menampilkan hasil
	fmt.Println("Hasil perolehan suara:")
	for _, p := range partaiList {
		fmt.Printf("%d(%d) ", p.Nama, p.Suara)
	}
	fmt.Println()
}
