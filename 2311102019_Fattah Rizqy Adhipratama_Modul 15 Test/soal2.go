package main

import (
	"fmt"
)

// Konstanta dan tipe data
const nMax = 51

type Mahasiswa struct {
	NIM_2311102019   string
	Nama  string
	Nilai int
}

type ArrayMahasiswa [nMax]Mahasiswa

// Fungsi untuk mencari nilai pertama dari mahasiswa dengan NIM tertentu
func cariNilaiPertama(array ArrayMahasiswa, n int, nim_2311102019 string) int {
	for i := 0; i < n; i++ {
		if array[i].NIM_2311102019 == nim_2311102019 {
			return array[i].Nilai
		}
	}
	return -1 // Nilai tidak ditemukan
}

// Fungsi untuk mencari nilai terakhir dari mahasiswa dengan NIM tertentu
func cariNilaiTerakhir(array ArrayMahasiswa, n int, nim_2311102019 string) int {
	for i := n - 1; i >= 0; i-- {
		if array[i].NIM_2311102019 == nim_2311102019 {
			return array[i].Nilai
		}
	}
	return -1 // Nilai tidak ditemukan
}

func main() {
	var arrayMahasiswa ArrayMahasiswa
	var n int

	fmt.Print("Masukkan jumlah data mahasiswa (maksimum 51): ")
	fmt.Scan(&n)

	if n > nMax || n <= 0 {
		fmt.Println("Jumlah data melebihi batas atau tidak valid.")
		return
	}

	// Input data mahasiswa
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data mahasiswa ke-%d\n", i+1)
		fmt.Print("NIM_2311102019: ")
		fmt.Scan(&arrayMahasiswa[i].NIM_2311102019)
		fmt.Print("Nama: ")
		fmt.Scan(&arrayMahasiswa[i].Nama)
		fmt.Print("Nilai: ")
		fmt.Scan(&arrayMahasiswa[i].Nilai)
	}

	// Input NIM untuk pencarian
	var nim_2311102019 string
	fmt.Print("Masukkan NIM_2311102019 untuk mencari nilai: ")
	fmt.Scan(&nim_2311102019)

	// Mencari nilai pertama dan terakhir
	nilaiPertama := cariNilaiPertama(arrayMahasiswa, n, nim_2311102019)
	nilaiTerakhir := cariNilaiTerakhir(arrayMahasiswa, n, nim_2311102019)

	// Menampilkan hasil
	if nilaiPertama != -1 && nilaiTerakhir != -1 {
		fmt.Printf("Nilai pertama mahasiswa dengan NIM %s: %d\n", nim_2311102019, nilaiPertama)
		fmt.Printf("Nilai terakhir mahasiswa dengan NIM %s: %d\n", nim_2311102019, nilaiTerakhir)
	} else {
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan.\n", nim_2311102019)
	}
}