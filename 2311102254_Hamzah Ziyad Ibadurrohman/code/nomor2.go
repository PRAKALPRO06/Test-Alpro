package main

import (
	"fmt"
)

// Struct untuk merepresentasikan data mahasiswa
type Mahasiswa struct {
	NIM   string
	Nama  string
	Nilai int
}

// Fungsi untuk mencari nilai pertama mahasiswa berdasarkan NIM
func cariNilaiPertama(data []Mahasiswa, targetNIM string) int {
	for _, m := range data {
		if m.NIM == targetNIM {
			return m.Nilai
		}
	}
	return -1 // jika tidak ketemu
}

// Fungsi untuk mencari nilai terbesar mahasiswa berdasarkan NIM mahasiswa
func cariNilaiTerbesar(data []Mahasiswa, targetNIM string) int {
	maxNilai := -1
	for _, m := range data {
		if m.NIM == targetNIM {
			if m.Nilai > maxNilai {
				maxNilai = m.Nilai
			}
		}
	}
	return maxNilai
}

// Fungsi untuk menerima masukan data mahasiswa
func inputData() []Mahasiswa {
	var n int
	fmt.Print("jumlah data mahasiswa: ")
	fmt.Scan(&n)

	data := make([]Mahasiswa, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data mahasiswa ke-%d:\n", i+1)
		fmt.Print("NIM: ")
		fmt.Scan(&data[i].NIM)
		fmt.Print("Nama: ")
		fmt.Scan(&data[i].Nama)
		fmt.Print("Nilai: ")
		fmt.Scan(&data[i].Nilai)
	}

	return data
}

// Fungsi untuk menampilkan hasil pencarian
func tampilkanHasil(nilaiPertama, nilaiTerbesar int) {
	if nilaiPertama == -1 {
		fmt.Println("Data tidak ditemukan.")
	} else {
		fmt.Printf("Nilai pertama: %d\n", nilaiPertama)
	}

	if nilaiTerbesar == -1 {
		fmt.Println("Data tidak ditemukan.")
	} else {
		fmt.Printf("Nilai terbesar: %d\n", nilaiTerbesar)
	}
}

func main() {

	data := inputData()

	var NIM2311102254 string
	fmt.Print("Masukkan NIM yang ingin dicari: ")
	fmt.Scan(&NIM2311102254)

	nilaiPertama := cariNilaiPertama(data, NIM2311102254)
	nilaiTerbesar := cariNilaiTerbesar(data, NIM2311102254)

	tampilkanHasil(nilaiPertama, nilaiTerbesar)
}
