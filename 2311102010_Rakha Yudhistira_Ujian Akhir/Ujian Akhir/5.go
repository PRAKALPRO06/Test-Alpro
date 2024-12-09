package main

import (
	"fmt"
	"sort"
)

// membuat struct untuk partai
type Partai struct {
	nama  int
	Suara int
}

func main() {
	// menentukan nilai maksimalnya
	const NMAX = 1000000

	// membuat array dengan besar nilai maksimal yang ditentukan
	suara_2311102010 := make([]int, NMAX+1) 

	for {
		var input int
		fmt.Scan(&input)
		if input == -1 {
			break
		}
		if input >= 1 && input <= NMAX {
			suara_2311102010[input]++
		}
	}

	// array partai beserta suaranya
	var partai []Partai
	for i := 1; i <= NMAX; i++ {
		if suara_2311102010[i] > 0 {
			partai = append(partai, Partai{nama: i, Suara: suara_2311102010[i]})
		}
	}

	// mengurutkan besar suara partai  
	sort.Slice(partai, func(i, j int) bool {
		if partai[i].Suara == partai[j].Suara {
			return partai[i].nama < partai[j].nama
		}
		return partai[i].Suara > partai[j].Suara
	})

	// menampilkan hasil stelah dihitung
	for _, p := range partai {
		fmt.Printf("%d(%d) ", p.nama, p.Suara)
	}
}
