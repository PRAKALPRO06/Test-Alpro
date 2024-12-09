package main

import (
	"fmt"
)

type Mahasiswa struct {
	NIM   string
	Nama  string
	Nilai int
}

type ArrayMahasiswa233 [51]Mahasiswa

func main() {
	var (
		dataMahasiswa ArrayMahasiswa233
		N             int
	)

	fmt.Print("Masukkan jumlah data mahasiswa (N): ")
	fmt.Scan(&N)

	// Input data mahasiswa
	for i := 0; i < N; i++ {
		fmt.Printf("\nData Mahasiswa ke-%d:\n", i+1)
		fmt.Print("NIM: ")
		fmt.Scan(&dataMahasiswa[i].NIM)
		fmt.Print("Nama: ")
		fmt.Scan(&dataMahasiswa[i].Nama)
		fmt.Print("Nilai: ")
		fmt.Scan(&dataMahasiswa[i].Nilai)
	}

	// Menu pilihan
	var pilihan int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Cari nilai pertama mahasiswa berdasarkan NIM")
		fmt.Println("2. Cari nilai terbesar mahasiswa berdasarkan NIM")
		fmt.Println("3. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			var nimCari string
			fmt.Print("Masukkan NIM mahasiswa: ")
			fmt.Scan(&nimCari)
			nilaiPertama := cariNilaiPertama(dataMahasiswa, N, nimCari)
			if nilaiPertama != -1 {
				fmt.Printf("Nilai pertama mahasiswa dengan NIM %s adalah %d\n", nimCari, nilaiPertama)
			} else {
				fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
			}

		case 2:
			var nimCari string
			fmt.Print("Masukkan NIM mahasiswa: ")
			fmt.Scan(&nimCari)
			nilaiTerbesar := cariNilaiTerbesar(dataMahasiswa, N, nimCari)
			if nilaiTerbesar != -1 {
				fmt.Printf("Nilai terbesar mahasiswa dengan NIM %s adalah %d\n", nimCari, nilaiTerbesar)
			} else {
				fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
			}

		case 3:
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func cariNilaiPertama(data ArrayMahasiswa233, N int, nim string) int {
	for i := 0; i < N; i++ {
		if data[i].NIM == nim {
			return data[i].Nilai
		}
	}
	return -1 // Tidak ditemukan
}

func cariNilaiTerbesar(data ArrayMahasiswa233, N int, nim string) int {
	maxNilai := -1
	found := false
	for i := 0; i < N; i++ {
		if data[i].NIM == nim {
			found = true
			if data[i].Nilai > maxNilai {
				maxNilai = data[i].Nilai
			}
		}
	}
	if found {
		return maxNilai
	}
	return -1 // Tidak ditemukan
}
