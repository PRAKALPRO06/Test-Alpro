package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const NMAX = 1000000

// Struct untuk menyimpan nama partai dan jumlah suara
type Partai struct {
	nama_2311102241 int
	suara           int
}

// Fungsi untuk mencari indeks partai di dalam array
func posisi_2311102241(tabPartai []Partai, n int, nama_2311102241 int) int {
	for i := 0; i < n; i++ {
		if tabPartai[i].nama_2311102241 == nama_2311102241 {
			return i
		}
	}
	return -1
}

func main() {
	var partai []Partai // Array untuk menyimpan data partai
	var n int           // Jumlah partai yang ditemukan

	// Membaca input pengguna
	fmt.Println("Masukkan data suara (akhiri dengan -1):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Memproses input
	data := strings.Split(input, " ")
	for _, val := range data {
		suara, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println("Input tidak valid:", val)
			return
		}
		if suara == -1 {
			break
		}
		idx := posisi_2311102241(partai, n, suara)
		if idx == -1 {
			// Jika partai belum ada di array, tambahkan partai baru
			partai = append(partai, Partai{nama_2311102241: suara, suara: 1})
			n++
		} else {
			// Jika partai sudah ada, tambahkan jumlah suaranya
			partai[idx].suara++
		}
	}

	// Sort array berdasarkan suara (descending) dan nama_2311102241 (ascending jika suara sama)
	sort.Slice(partai, func(i, j int) bool {
		if partai[i].suara == partai[j].suara {
			return partai[i].nama_2311102241 < partai[j].nama_2311102241
		}
		return partai[i].suara > partai[j].suara
	})

	// Output hasil
	fmt.Println("Hasil perolehan suara:")
	for _, p := range partai {
		fmt.Printf("%d(%d) ", p.nama_2311102241, p.suara)
	}
	fmt.Println()
}
