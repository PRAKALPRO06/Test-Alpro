package main

import (
	"fmt"
	"sort"
)

// Struct untuk menyimpan nama partai dan jumlah suara
type Partai struct {
	Nama_2311102278  int
	Suara int
}

// Fungsi untuk menghitung perolehan suara setiap partai
func hitungSuara(input []int) []Partai {
	partaiMap := make(map[int]int)

	// Menghitung jumlah suara setiap partai
	for _, suara := range input {
		partaiMap[suara]++
	}

	// Memindahkan data dari map ke slice
	var partaiList []Partai
	for nama_2311102278, suara := range partaiMap {
		partaiList = append(partaiList, Partai{Nama_2311102278: nama_2311102278, Suara: suara})
	}

	// Mengurutkan partai berdasarkan jumlah suara (descending)
	sort.Slice(partaiList, func(i, j int) bool {
		if partaiList[i].Suara == partaiList[j].Suara {
			return partaiList[i].Nama_2311102278 < partaiList[j].Nama_2311102278 // Jika sama, urutkan berdasarkan nama partai (ascending)
		}
		return partaiList[i].Suara > partaiList[j].Suara
	})

	return partaiList
}

func main() {
	var input []int
	var suara int

	fmt.Println("Masukkan suara partai (akhiri dengan -1):")
	for {
		fmt.Scan(&suara)
		if suara == -1 {
			break
		}
		input = append(input, suara)
	}

	// Menghitung hasil suara
	hasil := hitungSuara(input)

		// Menampilkan hasil
		fmt.Println("Hasil perolehan suara:")
		for _, partai := range hasil {
			fmt.Printf("%d(%d) ", partai.Nama_2311102278, partai.Suara)
		}
		fmt.Println()
	}
