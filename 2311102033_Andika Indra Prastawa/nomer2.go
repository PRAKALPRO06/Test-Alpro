package main

import "fmt"

// Maksimum jumlah data Mahasiswa
const nMax = 51

// Menyimpan data Mahasiswa
type Mahasiswa_2311102033 struct {
	NIM   string
	nama  string
	nilai int
}

// Untuk menyimpan data Mahasiswa
type arrayMahasiswa_2311102033 [nMax]Mahasiswa_2311102033

// Menerima input data Mahasiswa
func inputMahasiswa_2311102033(T *arrayMahasiswa_2311102033, N *int) {
	fmt.Print("Masukkan jumlah data Mahasiswa: ")
	fmt.Scan(N)
	for i := 0; i < *N; i++ {
		fmt.Printf("Data Mahasiswa ke-%d:\n", i+1)
		fmt.Print("Masukkan nim  mahasiswa: ")
		fmt.Scan(&T[i].NIM)
		fmt.Print("Masukkan nama mahasiswa: ")
		fmt.Scan(&T[i].nama)
		fmt.Print("Masukkan nilai mahasiswa: ")
		fmt.Scan(&T[i].nilai)
	}
}

//nilai pertama Mahasiswa berdasarkan dengan NIM
func carinilaipertama_2311102033(T arrayMahasiswa_2311102033, N int, nim string) int {
	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			return T[i].nilai
		}
	}
	return -1 // Jika tidak ditemukan
}

// Mencari nilai terbesar mahasiswa berdasarkan dengan NIM
func cariNilaiTerbesar(T arrayMahasiswa_2311102033, N int, nim string) int {
	maxNilai := -1
	found := false
	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
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
	var dataMahasiswa_2311102033 arrayMahasiswa_2311102033
	var jumlahData int
	var nim string

	// Input data Mahasiswa
	inputMahasiswa_2311102033(&dataMahasiswa_2311102033, &jumlahData)

	// Cari nilai pertama berdasarkan NIM
	fmt.Print("Masukkan NIM untuk mencari nilai mahasiswa pertama: ")
	fmt.Scan(&nim)
	nilaiPertama := carinilaipertama_2311102033(dataMahasiswa_2311102033, jumlahData, nim)
	if nilaiPertama != -1 {
		fmt.Printf("Nilai pertama untuk NIM %s adalah %d\n", nim, nilaiPertama)
	} else {
		fmt.Printf("Data dengan NIM %s tidak ditemukan\n", nim)
	}

	// Cari nilai terbesar berdasarkan NIM
	fmt.Print("Masukkan NIM untuk mencari nilai terbesar: ")
	fmt.Scan(&nim)
	nilaiTerbesar := cariNilaiTerbesar(dataMahasiswa_2311102033, jumlahData, nim)
	if nilaiTerbesar != -1 {
		fmt.Printf("Nilai terbesar untuk NIM %s adalah %d\n", nim, nilaiTerbesar)
	} else {
		fmt.Printf("Data dengan NIM %s tidak ditemukan\n", nim)
	}
}
