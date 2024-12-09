package main

import (
	"fmt"
)

// menentukan maksimal jumlah mahasiswa
const nMax = 51

// membuat struct mahasiswa
type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

// membuat array mahasiswa
type arrayMahasiswa [nMax]mahasiswa

// fungsi untuk menginputkan mahasiswa ke array mahasiswa
func inputMahasiswa_2311102010(T *arrayMahasiswa, N *int) {
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(N)

	for i := 0; i < *N; i++ {
		fmt.Printf("Masukkan data mahasiswa ke-%d:\n", i+1)
		fmt.Print("NIM: ")
		fmt.Scan(&T[i].NIM)
		fmt.Print("Nama: ")
		fmt.Scan(&T[i].nama)
		fmt.Print("Nilai: ")
		fmt.Scan(&T[i].nilai)
	}
}

// fungsi untuk menampilkan array mahasiswa
func printMahasiswa(T arrayMahasiswa, N int) {
	fmt.Println("Data mahasiswa:")
	for i := 0; i < N; i++ {
		fmt.Printf("%d. NIM: %s, Nama: %s, Nilai: %d\n", i+1, T[i].NIM, T[i].nama, T[i].nilai)
	}
}

// fungsi untuk menambahkan nilai baru berdasarkan nim
func tambahNilai(T *arrayMahasiswa, N int, nim string, nilaiBaru int) {
	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			T[i].nilai = nilaiBaru
			fmt.Printf("Nilai mahasiswa dengan NIM %s berhasil diupdate menjadi %d.\n", nim, nilaiBaru)
			return
		}
	}
	fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan.\n", nim)
}

// fungsi untuk mencari nilai berdasarkan nim
func cariNilai(T arrayMahasiswa, N int, nim string) {
	for i := 0; i < N; i++ {
		if T[i].NIM == nim {
			fmt.Printf("Mahasiswa dengan NIM %s memiliki nilai %d.\n", nim, T[i].nilai)
			return
		}
	}
	fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan.\n", nim)
}

func main() {
	var dataMahasiswa arrayMahasiswa
	var jumlahMahasiswa int

	inputMahasiswa_2311102010(&dataMahasiswa, &jumlahMahasiswa)
	printMahasiswa(dataMahasiswa, jumlahMahasiswa)

	var nim string
	var nilaiBaru int
	fmt.Print("Masukkan NIM mahasiswa untuk menambahkan nilai baru: ")
	fmt.Scan(&nim)
	fmt.Print("Masukkan nilai baru: ")
	fmt.Scan(&nilaiBaru)
	tambahNilai(&dataMahasiswa, jumlahMahasiswa, nim, nilaiBaru)

	fmt.Print("Masukkan NIM mahasiswa untuk mencari nilai: ")
	fmt.Scan(&nim)
	cariNilai(dataMahasiswa, jumlahMahasiswa, nim)
}