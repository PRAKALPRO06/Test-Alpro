package main

import (
	"fmt"
	"strings"
)

// Struck untuk menyimpan data mahasiswa
type mahasiswa_2311102327 struct {
	nim   string
	nama  string
	nilai int
}

// Fungsi untuk menambahkan data mahasiswa
func inputDataMahasiswa(jumlah int) []mahasiswa_2311102327 {
	data := make([]mahasiswa_2311102327, jumlah)
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Mahasiswa ke-%d:\n", i+1)
		fmt.Print("NIM: ")
		fmt.Scan(&data[i].nim)
		fmt.Print("Nama: ")
		fmt.Scan(&data[i].nama)
		fmt.Print("Nilai: ")
		fmt.Scan(&data[i].nilai)
		fmt.Println()
	}
	return data
}

// Fungsi untuk menampilkan data mahasiswa
func cariMahasiswa(nim string, data []mahasiswa_2311102327) *mahasiswa_2311102327 {
	for i, mhs := range data {
		if strings.EqualFold(mhs.nim, nim) {
			return &data[i]
		}
	}
	return nil
}

// Fungsi untuk hitung rata-rata nilai dengan perulangan
func rataRataNilai(data []mahasiswa_2311102327) float64 {
	total := 0
	for _, mhs := range data {
		total += mhs.nilai
	}
	return float64(total) / float64(len(data))
}

func main() {
	var jumlah int
	fmt.Print("Input jumlah mahasiswa nya: ")
	fmt.Scan(&jumlah)

	mahasiswaList := inputDataMahasiswa(jumlah)

	var nimCari string
	fmt.Print("\nMasukkan NIM untuk mencari mahasiswa: ")
	fmt.Scan(&nimCari)
	hasilCari := cariMahasiswa(nimCari, mahasiswaList)
	if hasilCari != nil {
		fmt.Printf("Mahasiswa ditemukan: NIM=%s, Nama=%s, Nilai=%d\n", hasilCari.nim, hasilCari.nama, hasilCari.nilai)
	} else {
		fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
	}

	rataNilai := rataRataNilai(mahasiswaList)
	fmt.Printf("\nRata-rata nilai mahasiswa: %.2f\n", rataNilai)
}