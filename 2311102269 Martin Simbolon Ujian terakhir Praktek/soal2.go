// Martin Simbolon 2311102269

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mahasiswa struct {
	NIM   string
	Nama  string
	Nilai []int
}

func main() {
	// Menerima masukan jumlah data mahasiswa
	fmt.Print("Masukkan jumlah data mahasiswa: ")
	reader := bufio.NewReader(os.Stdin)
	numInputStr, _ := reader.ReadString('\n')
	N, _ := strconv.Atoi(strings.TrimSpace(numInputStr))

	// Buat slice untuk menyimpan data mahasiswa
	dataMahasiswa := make([]Mahasiswa, 0, N)

	// Menerima data mahasiswa
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan NIM mahasiswa %d: ", i+1)
		nimInput, _ := reader.ReadString('\n')
		nim := strings.TrimSpace(nimInput)

		fmt.Printf("Masukkan nama mahasiswa %d: ", i+1)
		namaInput, _ := reader.ReadString('\n')
		nama := strings.TrimSpace(namaInput)

		fmt.Printf("Masukkan %d nilai mahasiswa %d (pisahkan dengan spasi): ", N, i+1)
		nilaiInput, _ := reader.ReadString('\n')
		nilaiStr := strings.Split(strings.TrimSpace(nilaiInput), " ")
		nilai := make([]int, 0, N)
		for _, v := range nilaiStr {
			v, _ := strconv.Atoi(v)
			nilai = append(nilai, v)
		}

		dataMahasiswa = append(dataMahasiswa, Mahasiswa{
			NIM:   nim,
			Nama:  nama,
			Nilai: nilai,
		})
	}

	// Mencari mahasiswa dengan nilai tertinggi di awal
	mahasiswaAwal := findHighestFirstValue(dataMahasiswa)
	fmt.Printf("Mahasiswa dengan nilai tertinggi di awal: %s (%s)\n", mahasiswaAwal.Nama, mahasiswaAwal.NIM)

	// Mencari mahasiswa dengan nilai tertinggi di akhir
	mahasiswaAkhir := findHighestLastValue(dataMahasiswa)
	fmt.Printf("Mahasiswa dengan nilai tertinggi di akhir: %s (%s)\n", mahasiswaAkhir.Nama, mahasiswaAkhir.NIM)
}

func findHighestFirstValue(data []Mahasiswa) Mahasiswa {
	var highest Mahasiswa
	for _, m := range data {
		if len(m.Nilai) > 0 && (len(highest.Nilai) == 0 || m.Nilai[0] > highest.Nilai[0]) {
			highest = m
		}
	}
	return highest
}

func findHighestLastValue(data []Mahasiswa) Mahasiswa {
	var highest Mahasiswa
	for _, m := range data {
		if len(m.Nilai) > 0 && (len(highest.Nilai) == 0 || m.Nilai[len(m.Nilai)-1] > highest.Nilai[len(highest.Nilai)-1]) {
			highest = m
		}
	}
	return highest
}