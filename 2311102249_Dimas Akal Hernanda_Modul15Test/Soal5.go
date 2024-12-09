package main

import (
	"fmt"
	"sort"
)

const maxPartai = 1000000

type Partai struct {
	Nama  int
	Suara int
}

func main() {
	var suara [maxPartai + 1]int
	var input int

	fmt.Println("Masukkan suara partai (akhiri dengan -1)_2311102249:")
	for {
		fmt.Scan(&input)
		if input == -1 {
			break
		}
		if input >= 1 && input <= maxPartai {
			suara[input]++
		}
	}

	var hasil []Partai
	for i := 1; i <= maxPartai; i++ {
		if suara[i] > 0 {
			hasil = append(hasil, Partai{Nama: i, Suara: suara[i]})
		}
	}

	sort.Slice(hasil, func(i, j int) bool {
		return hasil[i].Suara > hasil[j].Suara
	})

	fmt.Println("Keluaran:")
	for _, p := range hasil {
		fmt.Printf("%d(%d) ", p.Nama, p.Suara)
	}
	fmt.Println()
}
