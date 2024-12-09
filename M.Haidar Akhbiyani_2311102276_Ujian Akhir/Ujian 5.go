package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Struct untuk menyimpan partai dan suaranya
type Partai struct {
	ID    int
	Suara int
}

// Program ini dibuat oleh [NIM 2311102276]
func main() {
	fmt.Println("Program Pemrosesan Suara Partai - Dibuat oleh NIM 2311102276")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan data (akhiri dengan -1):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Memproses input
	data := strings.Split(input, " ")
	counter := make(map[int]int)

	// Menghitung suara per partai
	for _, s := range data {
		num, _ := strconv.Atoi(s)
		if num == -1 {
			break
		}
		counter[num]++
	}

	// Memasukkan data ke dalam slice untuk diurutkan
	var partaiList []Partai
	for id, suara := range counter {
		partaiList = append(partaiList, Partai{ID: id, Suara: suara})
	}

	// Mengurutkan berdasarkan jumlah suara (descending), lalu ID partai (ascending)
	sort.Slice(partaiList, func(i, j int) bool {
		if partaiList[i].Suara == partaiList[j].Suara {
			return partaiList[i].ID < partaiList[j].ID
		}
		return partaiList[i].Suara > partaiList[j].Suara
	})

	// Menampilkan output
	fmt.Println("Hasil perhitungan suara:")
	for _, p := range partaiList {
		fmt.Printf("%d(%d) ", p.ID, p.Suara)
	}
	fmt.Println() // Tambahkan baris baru setelah output
}
