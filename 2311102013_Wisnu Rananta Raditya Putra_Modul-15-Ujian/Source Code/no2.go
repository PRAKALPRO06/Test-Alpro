package main

import "fmt"

// Konstanta maksimum jumlah data mahasiswa
const nMax = 51

// Struct untuk menyimpan data mahasiswa
type mahasiswa struct {
	NIM_2311102013   string
	nama  string
	nilai int
}

// Array untuk menyimpan data mahasiswa
type arrayMahasiswa [nMax]mahasiswa

// Fungsi untuk menerima input data mahasiswa
func inputMahasiswa(T *arrayMahasiswa, N *int) {
	fmt.Print("Masukkan jumlah data mahasiswa: ")
	fmt.Scan(N)
	for i := 0; i < *N; i++ {
		fmt.Printf("Data mahasiswa ke-%d:\n", i+1)
		fmt.Print("Masukkan NIM: ")
		fmt.Scan(&T[i].NIM_2311102013)
		fmt.Print("Masukkan nama: ")
		fmt.Scan(&T[i].nama)
		fmt.Print("Masukkan nilai: ")
		fmt.Scan(&T[i].nilai)
	}
}

// Fungsi untuk mencari nilai pertama seorang mahasiswa berdasarkan NIM
func cariNilaiPertama(T arrayMahasiswa, N int, nim string) int {
	for i := 0; i < N; i++ {
		if T[i].NIM_2311102013 == nim {
			return T[i].nilai
		}
	}
	return -1 // Jika tidak ditemukan
}

// Fungsi untuk mencari nilai terbesar seorang mahasiswa berdasarkan NIM
func cariNilaiTerbesar(T arrayMahasiswa, N int, nim string) int {
	maxNilai := -1
	found := false
	for i := 0; i < N; i++ {
		if T[i].NIM_2311102013 == nim {
			found = true
			if T[i].nilai > maxNilai {
				maxNilai = T[i].nilai
			}
		}
	}
	if found {
		return maxNilai
	}
	return -1 // Jika tidak ditemukan
}

func main() {
	var dataMahasiswa arrayMahasiswa
	var jumlahData int
	var nim string

	// Input data mahasiswa
	inputMahasiswa(&dataMahasiswa, &jumlahData)

	// Cari nilai pertama berdasarkan NIM
	fmt.Print("Masukkan NIM untuk mencari nilai pertama: ")
	fmt.Scan(&nim)
	nilaiPertama := cariNilaiPertama(dataMahasiswa, jumlahData, nim)
	if nilaiPertama != -1 {
		fmt.Printf("Nilai pertama untuk NIM %s adalah %d\n", nim, nilaiPertama)
	} else {
		fmt.Printf("Data dengan NIM %s tidak ditemukan\n", nim)
	}

	// Cari nilai terbesar berdasarkan NIM
	fmt.Print("Masukkan NIM untuk mencari nilai terbesar: ")
	fmt.Scan(&nim)
	nilaiTerbesar := cariNilaiTerbesar(dataMahasiswa, jumlahData, nim)
	if nilaiTerbesar != -1 {
		fmt.Printf("Nilai terbesar untuk NIM %s adalah %d\n", nim, nilaiTerbesar)
	} else {
		fmt.Printf("Data dengan NIM %s tidak ditemukan\n", nim)
	}
}
