package main

import "fmt"

// definisi jumlah maks mahasiswa
const nMax = 51

// Struct untuk menyimpan data mahasiswa
type Mahasiswa struct {
	NIM_2311102278   string
	Nama  string
	Nilai int
}

// Array untuk menyimpan data mahasiswa
var arrayMahasiswa [nMax]Mahasiswa

// Fungsi untuk memasukkan data mahasiswa ke dalam array
func inputDataMahasiswa(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data mahasiswa ke-%d\n", i+1)
		fmt.Print("NIM: ")
		fmt.Scan(&arrayMahasiswa[i].NIM_2311102278)
		fmt.Print("Nama: ")
		fmt.Scan(&arrayMahasiswa[i].Nama)
		fmt.Print("Nilai: ")
		fmt.Scan(&arrayMahasiswa[i].Nilai)
	}
}

// Fungsi untuk mencari nilai pertama berdasarkan NIM
func cariNilaiPertama(n int, nim string) int {
	for i := 0; i < n; i++ {
		if arrayMahasiswa[i].NIM_2311102278 == nim {
			return arrayMahasiswa[i].Nilai
		}
	}
	return -1 // Mengembalikan -1 jika NIM tidak ditemukan
}

// Fungsi untuk mencari nilai terbesar berdasarkan NIM
func cariNilaiTerbesar(n int, nim string) int {
	maxNilai := -1
	for i := 0; i < n; i++ {
		if arrayMahasiswa[i].NIM_2311102278 == nim && arrayMahasiswa[i].Nilai > maxNilai {
			maxNilai = arrayMahasiswa[i].Nilai
		}
	}
	return maxNilai
}

// Fungsi utama
func main() {
	var n int
	var nim string

	// Input jumlah mahasiswa
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&n)

	// Memastikan jumlah mahasiswa tidak melebihi batas
	if n > nMax {
		fmt.Println("Jumlah mahasiswa melebihi batas!")
		return
	}

	// Input data mahasiswa
	inputDataMahasiswa(n)

	// Mencari nilai pertama berdasarkan NIM
	fmt.Print("Masukkan NIM untuk mencari nilai pertama: ")
	fmt.Scan(&nim)
	nilaiPertama := cariNilaiPertama(n, nim)
	if nilaiPertama == -1 {
		fmt.Println("NIM tidak ditemukan.")
	} else {
		fmt.Printf("Nilai pertama mahasiswa dengan NIM %s adalah %d\n", nim, nilaiPertama)
	}

	// Mencari nilai terbesar berdasarkan NIM
	fmt.Print("Masukkan NIM untuk mencari nilai terbesar: ")
	fmt.Scan(&nim)
	nilaiTerbesar := cariNilaiTerbesar(n, nim)
	if nilaiTerbesar == -1 {
		fmt.Println("NIM tidak ditemukan.")
	} else {
		fmt.Printf("Nilai terbesar mahasiswa dengan NIM %s adalah %d\n", nim, nilaiTerbesar)
	}
}
