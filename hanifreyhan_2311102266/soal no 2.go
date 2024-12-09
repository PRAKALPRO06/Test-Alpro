package main

import (
	"fmt"
)
type Mahasiswa struct {
	NIM   string
	Nama  string
	Nilai int
}
const nMax = 5
type ArrayMahasiswa [nMax]Mahasiswa
func main() {
	var dataMahasiswa ArrayMahasiswa
	var n int

	// Input jumlah data mahasiswa
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scanln(&n)

	// Validasi jumlah mahasiswa
	if n <= 0 || n > nMax {
		fmt.Printf("Jumlah mahasiswa harus antara 1 hingga %d\n", nMax)
		return
	}

	// Input data mahasiswa
	for i := 0; i < n; i++ {
		fmt.Printf("\nMasukkan data mahasiswa ke-%d:\n", i+1)
		fmt.Print("NIM: ")
		fmt.Scan(&dataMahasiswa[i].NIM)
		fmt.Print("Nama: ")
		fmt.Scan(&dataMahasiswa[i].Nama)
		fmt.Print("Nilai: ")
		fmt.Scan(&dataMahasiswa[i].Nilai)
	}

	// Input NIM untuk pencarian
	var cariNIM string
	fmt.Print("\nMasukkan NIM untuk pencarian: ")
	fmt.Scan(&cariNIM)

	// Cari nilai pertama mahasiswa berdasarkan NIM
	nilaiPertama := cariNilaiPertama(dataMahasiswa[:n], cariNIM)
	if nilaiPertama != -1 {
		fmt.Printf("Nilai awal NIM %s adalah %d\n", cariNIM, nilaiPertama)
	} else {
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan\n", cariNIM)
	}

	// Cari nilai terbesar mahasiswa berdasarkan NIM
	nilaiTerbesar := cariNilaiTerbesar(dataMahasiswa[:n], cariNIM)
	if nilaiTerbesar != -1 {
		fmt.Printf("Nilai terbesar mahasiswa dengan NIM %s adalah %d\n", cariNIM, nilaiTerbesar)
	} else {
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan\n", cariNIM)
	}
}

func cariNilaiPertama(data []Mahasiswa, nim string) int {
	for _, mhs := range data {
		if mhs.NIM == nim {
			return mhs.Nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(data []Mahasiswa, nim string) int {
	var maxNilai = -1
	for _, mhs := range data {
		if mhs.NIM == nim {
			if mhs.Nilai > maxNilai {
				maxNilai = mhs.Nilai
			}
		}
	}
	return maxNilai
}
