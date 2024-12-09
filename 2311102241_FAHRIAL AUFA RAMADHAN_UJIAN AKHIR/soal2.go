package main

import (
	"fmt"
)

const nMax int = 51

// Struct untuk menyimpan data mahasiswa
type Mahasiswa struct {
	NIM_2311102241 string
	Nama           string
	Nilai          int
}

// Array untuk menyimpan data mahasiswa
type ArrayMahasiswa [nMax]Mahasiswa

// Fungsi untuk menambahkan data mahasiswa ke dalam array
func input(arr *ArrayMahasiswa, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan NIM: ")
		fmt.Scan(&arr[i].NIM_2311102241)
		fmt.Printf("Masukkan Nama: ")
		fmt.Scan(&arr[i].Nama)
		fmt.Printf("Masukkan Nilai: ")
		fmt.Scan(&arr[i].Nilai)
	}
}

// Fungsi untuk mencari nilai pertama dari mahasiswa berdasarkan NIM_2311102241
func nilaiPertama_2311102241(arr *ArrayMahasiswa, n int, nim_2311102241 string) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM_2311102241 == nim_2311102241 {
			return arr[i].Nilai
		}
	}
	return -1 // Jika tidak ditemukan
}

// Fungsi untuk mencari nilai terbesar dari mahasiswa berdasarkan NIM_2311102241
func nilaiTerbesar_2311102241(arr *ArrayMahasiswa, n int, nim_2311102241 string) int {
	maxNilai := -1
	for i := 0; i < n; i++ {
		if arr[i].NIM_2311102241 == nim_2311102241 && arr[i].Nilai > maxNilai {
			maxNilai = arr[i].Nilai
		}
	}
	return maxNilai
}

func main() {
	var arr ArrayMahasiswa
	var n int
	var nim_2311102241 string

	fmt.Print("Masukkan jumlah data (N): ")
	fmt.Scan(&n)

	if n > nMax {
		fmt.Printf("Jumlah data mahasiswa tidak boleh lebih dari %d\n", nMax)
		return
	}

	input(&arr, n)

	fmt.Print("Masukkan NIM untuk mencari nilai pertama: ")
	fmt.Scan(&nim_2311102241)
	nilaiPertama := nilaiPertama_2311102241(&arr, n, nim_2311102241)
	if nilaiPertama != -1 {
		fmt.Printf("Nilai pertama mahasiswa dengan NIM %s adalah %d\n", nim_2311102241, nilaiPertama)
	} else {
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan\n", nim_2311102241)
	}

	fmt.Print("Masukkan NIM untuk mencari nilai terbesar: ")
	fmt.Scan(&nim_2311102241)
	nilaiTerbesar := nilaiTerbesar_2311102241(&arr, n, nim_2311102241)
	if nilaiTerbesar != -1 {
		fmt.Printf("Nilai terbesar mahasiswa dengan NIM %s adalah %d\n", nim_2311102241, nilaiTerbesar)
	} else {
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan\n", nim_2311102241)
	}
}
