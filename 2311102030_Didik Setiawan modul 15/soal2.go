package main

import (
	"fmt"
	"strconv"
	"strings"
)

const nMax = 51

// Definisi struct mahasiswa
type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

// Array untuk menyimpan data mahasiswa
type arrayMahasiswa [nMax]mahasiswa

// Fungsi untuk membaca data mahasiswa dari input
func inputMahasiswa(n int) arrayMahasiswa {
	var arr arrayMahasiswa

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data mahasiswa ke-%d (NIM Nama Nilai): ", i+1)

		data := strings.Fields()

		if len(data) < 3 {
			fmt.Println("Input tidak valid. Masukkan NIM, nama, dan nilai.")
			i--
			continue
		}

		nilai, err := strconv.Atoi(data[2])
		if err != nil {
			fmt.Println("Nilai harus berupa angka.")
			i--
			continue
		}

		arr[i] = mahasiswa{
			NIM:   data[0],
			nama:  data[1],
			nilai: nilai,
		}
	}
	return arr
}

// Fungsi untuk mencari nilai berdasarkan NIM
func cariNilai(arr arrayMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if arr[i].NIM == nim {
			return arr[i].nilai
		}
	}
	return -1 // Jika NIM tidak ditemukan
}

// Fungsi untuk mencari nilai tertinggi
func nilaiTertinggi(arr arrayMahasiswa, n int) mahasiswa {
	max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i].nilai > max.nilai {
			max = arr[i]
		}
	}
	return max
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&n)

	if n > nMax {
		fmt.Printf("Jumlah mahasiswa tidak boleh lebih dari %d.\n", nMax)
		return
	}

	// Input data mahasiswa
	mahasiswaArray := inputMahasiswa(n)

	// Mencari nilai berdasarkan NIM
	var nim string
	fmt.Print("Masukkan NIM untuk mencari nilai: ")
	fmt.Scan(&nim)
	nilai := cariNilai(mahasiswaArray, n, nim)
	if nilai != -1 {
		fmt.Printf("Nilai mahasiswa dengan NIM %s adalah %d.\n", nim, nilai)
	} else {
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan.\n", nim)
	}

	// Mencari mahasiswa dengan nilai tertinggi
	mahasiswaMax := nilaiTertinggi(mahasiswaArray, n)
	fmt.Printf("Mahasiswa dengan nilai tertinggi adalah %s (%s) dengan nilai %d.\n",
		mahasiswaMax.nama, mahasiswaMax.NIM, mahasiswaMax.nilai)
}
