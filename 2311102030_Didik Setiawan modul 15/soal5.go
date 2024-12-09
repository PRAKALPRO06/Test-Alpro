package main

import (
	"fmt"
	"sort"
)

// struct untuk partai_2311102030
type Partai_2311102030 struct {
	Nama  int
	Suara int
}

func main() {
	const NMAX = 1000000
	suara := make([]int, NMAX+1) //tipe tab partai_2311102030

	for {
		var input int
		fmt.Scan(&input)
		if input == -1 {
			break
		}
		if input >= 1 && input <= NMAX {
			suara[input]++
		}
	}

	// Kumpulkan partai_2311102030 dengan suara
	var partai_2311102030 []Partai_2311102030
	for i := 1; i <= NMAX; i++ {
		if suara[i] > 0 {
			partai_2311102030 = append(partai_2311102030, Partai_2311102030{Nama: i, Suara: suara[i]})
		}
	}

	// mengurutkan partai_2311102030 bedasarkan partai_2311102030 dan suara secara menurun
	sort.Slice(partai_2311102030, func(i, j int) bool {
		if partai_2311102030[i].Suara == partai_2311102030[j].Suara {
			return partai_2311102030[i].Nama < partai_2311102030[j].Nama
		}
		return partai_2311102030[i].Suara > partai_2311102030[j].Suara
	})

	// menampilkan hasil
	for _, p := range partai_2311102030 {
		fmt.Printf("%d(%d) ", p.Nama, p.Suara)
	}
}
