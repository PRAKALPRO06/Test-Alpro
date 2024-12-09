package main
import (
	"fmt"
)
const NMAX = 1000000
type Partai struct {
	nama  int
	suara int
}
type TabPartai [NMAX]Partai

// Fungsi untuk mencari indeks partai berdasarkan nama
func posisi(t TabPartai, n int, nama int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i
		}
	}
	return -1
}

// Fungsi untuk mengurutkan array partai secara descending berdasarkan jumlah suara, untuk memindahkan elemen yang lebih kecil menjadi ke kanan
func insertionSortDescending(t *TabPartai, n int) {
	for i := 1; i < n; i++ {
		key := t[i]
		j := i - 1

		for j >= 0 && t[j].suara < key.suara {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = key
	}
}

func main() {
	var t_2311102011 TabPartai
	var n int
	var input int
	n = 0

	for {
		fmt.Scan(&input)
		if input == -1 {
			break
		}

		pos := posisi(t_2311102011, n, input)

		if pos == -1 {
			t_2311102011[n].nama = input
			t_2311102011[n].suara = 1
			n++
		} else {
			t_2311102011[pos].suara++
		}
	}
	insertionSortDescending(&t_2311102011, n)

	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d) ", t_2311102011[i].nama, t_2311102011[i].suara)
	}
	fmt.Println()
}
