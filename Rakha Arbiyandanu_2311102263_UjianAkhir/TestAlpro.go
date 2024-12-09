package main

import (
	"fmt"
	"sort"
)

// Struct untuk menyimpan data partai (ID dan jumlah suara)
type Partai_2311102263 struct {
	ID    int
	Suara int
}

func main() {//fungsi utaa untuk menjalankan fungsi fungsi 
	
	const PENANDA_AKHIR = -1
	const JUMLAH_MAX_PARTAI = 1000000

	// Array untuk menyimpan jumlah suara tiap partai
	suara := make([]int, JUMLAH_MAX_PARTAI+1)

	var input int

	// Membaca input suara
	for {
		fmt.Scan(&input)
		if input == PENANDA_AKHIR {
			break
		}
		// Menambahkan suara ke partai terkait
		suara[input]++
	}


	var daftarPartai []Partai_2311102263
	for i := 1; i <= JUMLAH_MAX_PARTAI; i++ {
		if suara[i] > 0 {
			daftarPartai = append(daftarPartai, Partai_2311102263{ID: i, Suara: suara[i]})
		}
	}

	// Mengurutkan partai berdasarkan jumlah suara (descending)
	// Jika jumlah suara sama, diurutkan berdasarkan ID partai (ascending)
	sort.Slice(daftarPartai, func(i, j int) bool {
		if daftarPartai[i].Suara == daftarPartai[j].Suara {
			return daftarPartai[i].ID < daftarPartai[j].ID
		}
		return daftarPartai[i].Suara > daftarPartai[j].Suara
	})

	// Untuk Menampilkan hasil
	for _, partai := range daftarPartai {
		fmt.Printf("%d(%d) ", partai.ID, partai.Suara)
	}
	fmt.Println()
}
