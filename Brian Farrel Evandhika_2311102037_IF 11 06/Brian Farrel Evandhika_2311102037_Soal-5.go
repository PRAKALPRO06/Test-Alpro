// 2311102037 Brian Farrel Evandhika IF 11 06
package main

import (
	"fmt"
)

const NMAX = 1000000

type Partai struct {
	nama_2311102037 int
	suara           int
}

type tabPartai [NMAX]Partai

func main() { //fungsi utama untuk memanggil fungsi lainya
	var partaiList tabPartai
	var n int

	fmt.Println("Masukkan nama partai (akhiri dengan -1):") //memasukan input suara pencoblosan
	for {
		var input int
		fmt.Scan(&input)
		if input == -1 {
			break
		}
		pos := posisi(partaiList, n, input)
		if pos == -1 {
			partaiList[n] = Partai{nama_2311102037: input, suara: 1}
			n++
		} else {
			partaiList[pos].suara++
		}
	}

	for i := 1; i < n; i++ { //pengurutan dengan menggunakan insertion sort descending
		key := partaiList[i]
		j := i - 1
		for j >= 0 && partaiList[j].suara < key.suara {
			partaiList[j+1] = partaiList[j]
			j--
		}
		partaiList[j+1] = key
	}

	fmt.Println("Perolehan suara partai:") //menampilkan hasil
	for i := 0; i < n; i++ {
		fmt.Printf("Partai %d: %d suara\n", partaiList[i].nama_2311102037, partaiList[i].suara)
	}
}

func posisi(t tabPartai, n int, nama_2311102037 int) int { //fungsi posisi mencari posisi (indeks) dari elemen yang memiliki nilai nama_2311102037 pada array t
	for i := 0; i < n; i++ {
		if t[i].nama_2311102037 == nama_2311102037 {
			return i
		}
	}
	return -1
}
