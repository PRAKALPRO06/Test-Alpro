package main

import "fmt"

const NMAX = 1000000

// struct partai
type partai struct {
	nama  int
	suara int
}

func main() {
	// deklarasi variabel
	var p [NMAX]partai

	// lakukan proses input suara secara berulang di sini, simpan ke array p
	for i := 0; ; i++ {
		var nama, suara int
		_, err := fmt.Scanf("%d %d", &nama, &suara)
		if err != nil {
			break
		}
		p[i] = partai{nama, suara}
	}

	// lakukan proses pengurutan dengan insertion sort descending
	for i := 1; i < len(p); i++ {
		key := p[i]
		j := i - 1
		for j >= 0 && p[j].suara < key.suara {
			p[j+1] = p[j]
			j--
		}
		p[j+1] = key
	}

	// tampilkan array p
	for i := 0; i < len(p) && p[i].nama != 0; i++ {
		fmt.Printf("%d %d\n", p[i].nama, p[i].suara)
	}
}

func posisi(t []partai, n int, nama int) int {
	// mengembalikan indeks partai yang memiliki nama yang dicari pada array t
	// jika tidak ditemukan, kembalikan -1
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}
