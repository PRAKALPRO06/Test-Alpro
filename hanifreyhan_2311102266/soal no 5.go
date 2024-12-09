package main

import (
	"fmt"
	"sort"
)

type Party struct {
	Name  int
	Votes int
}

func main() {
	var input []int
	partyMap := make(map[int]int)
	fmt.Println("Masukkan angka partai diikuti dengan -1 untuk mengakhiri input:")

	// Baca input dari pengguna
	for {
		var num2311102266 int
		_, err := fmt.Scan(&num2311102266)
		if err != nil {
			fmt.Println("Terjadi kesalahan input. Pastikan hanya memasukkan angka.")
			return
		}
		if num2311102266 == -1 {
			break
		}
		input = append(input, num2311102266)
		partyMap[num2311102266]++
	}

	// Konversi map menjadi slice untuk diurutkan
	var parties []Party
	for name, votes := range partyMap {
		parties = append(parties, Party{Name: name, Votes: votes})
	}

	// Urutkan berdasarkan jumlah suara (descending), jika sama, berdasarkan nama partai (ascending)
	sort.Slice(parties, func(i, j int) bool {
		if parties[i].Votes == parties[j].Votes {
			return parties[i].Name < parties[j].Name
		}
		return parties[i].Votes > parties[j].Votes
	})

	// Cetak hasil
	fmt.Println("Hasil perolehan suara:")
	for _, party := range parties {
		fmt.Printf("%d(%d) ", party.Name, party.Votes)
	}
	fmt.Println()
}
