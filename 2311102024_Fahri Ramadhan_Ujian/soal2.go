package main

import "fmt"

const nMax = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

// Fungsi ini untuk menerima input data mahasiswa
func inputMahasiswa(T *arrayMahasiswa, N *int) {
	fmt.Print("Masukkan jumlah data mahasiswa: ")
	fmt.Scan(N)
	for i := 0; i < *N; i++ {
		fmt.Printf("Data mahasiswa ke-%d:\n", i+1)
		fmt.Print("Masukkan NIM: ")
		fmt.Scan(&T[i].NIM)
		fmt.Print("Masukkan nama: ")
		fmt.Scan(&T[i].nama)
		fmt.Print("Masukkan nilai: ")
		fmt.Scan(&T[i].nilai)
	}
}

// Fungsi ini untuk mencari nilai pertama seorang mahasiswa berdasarkan NIM
func cariNilaiPertama(T arrayMahasiswa, N int, nim string) int {
	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			return T[i].nilai
		}
	}
	return -1 // Jika tidak ditemukan
}

// Fungsi ini untuk mencari nilai terbesar seorang mahasiswa berdasarkan NIM
func cariNilaiTerbesar(T arrayMahasiswa, N int, nim string) int {
	maxNilai_2311102024 := -1
	found := false
	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			found = true
			if T[i].nilai > maxNilai_2311102024 {
				maxNilai_2311102024 = T[i].nilai
			}
		}
	}
	if found {
		return maxNilai_2311102024
	}
	return -1 // Jika tidak ditemukan
}

func main() {
	var dataMahasiswa arrayMahasiswa
	var jumlahData int
	var nim string

	inputMahasiswa(&dataMahasiswa, &jumlahData)

	fmt.Print("Masukkan NIM untuk mencari nilai pertama: ")
	fmt.Scan(&nim)
	nilaiPertama := cariNilaiPertama(dataMahasiswa, jumlahData, nim)
	if nilaiPertama != -1 {
		fmt.Printf("Nilai pertama untuk NIM %s adalah %d\n", nim, nilaiPertama)
	} else {
		fmt.Printf("Data dengan NIM %s tidak ditemukan\n", nim)
	}

	fmt.Print("Masukkan NIM untuk mencari nilai terbesar: ")
	fmt.Scan(&nim)
	nilaiTerbesar := cariNilaiTerbesar(dataMahasiswa, jumlahData, nim)
	if nilaiTerbesar != -1 {
		fmt.Printf("Nilai terbesar untuk NIM %s adalah %d\n", nim, nilaiTerbesar)
	} else {
		fmt.Printf("Data dengan NIM %s tidak ditemukan\n", nim)
	}
}