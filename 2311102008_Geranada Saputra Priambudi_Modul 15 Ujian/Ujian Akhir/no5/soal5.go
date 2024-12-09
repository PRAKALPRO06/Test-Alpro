// Geranada Saputra Priambudi
// 2311102008
// IF-11-06
package main

import "fmt"

const NMAX = 1000000

// struct partai
type partai struct {
	nama, suara int
}

// tipe tabPartai: array of partai dengan kapasitas NMAX
type tabPartai [NMAX]partai

func main() {
	// deklarasi variabel
	var p_2311102008 tabPartai
	var n, x, idx int
	// lakukan proses input suara secara berulang di sini, simpan kedalam array p_2311102008, sehingga terdapat array p_2311102008 yang berisi hasil peroleh suara n partai.
	n = 0
	fmt.Scan(&x)
	for x != -1 {
		idx = posisi(p_2311102008, n, x)
		if idx == -1 {
			p_2311102008[n].nama = x
			p_2311102008[n].suara = 1
			n++
		} else {
			p_2311102008[idx].suara++
		}
		fmt.Scan(&x)
	}
	// lakukan proses pengurutan dengan insertion sort descending berdasarkan jumlah suara yang diperoleh
	var pass, k int
	var temp partai
	for pass = 1; pass <= n-1; pass++ {
		k = pass
		temp = p_2311102008[k]
		for k > 0 && temp.suara > p_2311102008[k-1].suara {
			p_2311102008[k] = p_2311102008[k-1]
			k--
		}
		p_2311102008[k] = temp
	}
	// tampilkan array p_2311102008
	for k = 0; k < n; k++ {
		fmt.Printf("%v(%v) ", p_2311102008[k].nama, p_2311102008[k].suara)
	}
}
func posisi(t tabPartai, n int, nama int) int {
	/* mengembalikan indeks partai yang memiliki nama yang dicari pada array
	   t yang berisi n partai atau -1 apabila tidak ditemukan , gunakan
	   pencarian sekuensial */
	var i, ketemu int
	i = 0
	ketemu = -1
	for i < n && ketemu == -1 {
		if t[i].nama == nama {
			ketemu = i
		}
		i++
	}
	return ketemu
}
