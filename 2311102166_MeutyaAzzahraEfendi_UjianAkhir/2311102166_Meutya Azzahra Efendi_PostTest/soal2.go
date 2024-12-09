package main

import "fmt"

// Konstanta untuk ukuran maksimum array
const nMax = 51

// Struktur untuk data mahasiswa
type mahasiswa struct {
	NIM   string
	Nama  string
	Nilai int
}

// Array untuk menyimpan data mahasiswa
type arrayMahasiswa_2311102166 [nMax]mahasiswa

// Fungsi untuk menerima data mahasiswa
func inputData(N int, arr *arrayMahasiswa_2311102166) {
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan data mahasiswa ke-%d:\n", i+1)
		fmt.Print("NIM: ")
		fmt.Scan(&arr[i].NIM)
		fmt.Print("Nama: ")
		fmt.Scan(&arr[i].Nama)
		fmt.Print("Nilai: ")
		fmt.Scan(&arr[i].Nilai)
	}
}

// Fungsi untuk mencari data mahasiswa berdasarkan NIM
func cariMahasiswa(N int, arr *arrayMahasiswa_2311102166, nim string) int {
	for i := 0; i < N; i++ {
		if arr[i].NIM == nim {
			return i
		}
	}
	return -1 // Tidak ditemukan
}

// Fungsi untuk mencari nilai terbesar mahasiswa berdasarkan NIM 
func cariNilaiTerbesar(N int, arr *arrayMahasiswa_2311102166, nim string) int {
	idx := cariMahasiswa(N, arr, nim)
	if idx == -1 {
		fmt.Println("Mahasiswa dengan NIM", nim, "tidak ditemukan.")
		return -1
	}
	return arr[idx].Nilai
}

// Fungsi untuk menampilkan seluruh data mahasiswa
func tampilkanData(N int, arr *arrayMahasiswa_2311102166) {
	fmt.Println("Daftar Mahasiswa:")
	for i := 0; i < N; i++ {
		fmt.Printf("%d. NIM: %s, Nama: %s, Nilai: %d\n", i+1, arr[i].NIM, arr[i].Nama, arr[i].Nilai)
	}
}

func main() {
	var arr arrayMahasiswa_2311102166
	var N int

	// Input jumlah mahasiswa
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&N)

	// Input data mahasiswa
	inputData(N, &arr)

	// Menampilkan seluruh data mahasiswa
	tampilkanData(N, &arr)

	// Mencari mahasiswa berdasarkan NIM
	var nim string
	fmt.Print("Masukkan NIM untuk mencari mahasiswa: ")
	fmt.Scan(&nim)
	idx := cariMahasiswa(N, &arr, nim)
	if idx != -1 {
		fmt.Printf("Mahasiswa ditemukan: NIM: %s, Nama: %s, Nilai: %d\n", arr[idx].NIM, arr[idx].Nama, arr[idx].Nilai)
	} else {
		fmt.Println("Mahasiswa tidak ditemukan.")
	}

	// Mencari nilai terbesar berdasarkan NIM
	fmt.Print("Masukkan NIM untuk mencari nilai terbesar: ")
	fmt.Scan(&nim)
	nilai := cariNilaiTerbesar(N, &arr, nim)
	if nilai != -1 {
		fmt.Printf("Nilai terbesar untuk NIM %s adalah %d\n", nim, nilai)
	}
}
