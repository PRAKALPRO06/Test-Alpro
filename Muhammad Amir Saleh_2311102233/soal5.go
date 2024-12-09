package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama  int
	suara int
}

func main() {
	var (
		p      [NMAX]partai
		jumlah int
		input  int
	)

	fmt.Println("Masukkan data suara partai (akhiri dengan -1):")
	for {
		fmt.Scan(&input)
		if input == -1 {
			break
		}

		pos := posisi(p, jumlah, input)
		if pos == -1 {
			p[jumlah] = partai{nama: input, suara: 1}
			jumlah++
		} else {
			p[pos].suara++
		}
	}

	// Pengurutan dengan insertion sort secara descending
	for i := 1; i < jumlah; i++ {
		key := p[i]
		j := i - 1
		for j >= 0 && p[j].suara < key.suara {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = key
	}

	// Menampilkan hasil
	fmt.Println("Hasil perolehan suara:")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Partai %d (%d)\n", p[i].nama, p[i].suara)
	}
}

func posisi(t [NMAX]partai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}
