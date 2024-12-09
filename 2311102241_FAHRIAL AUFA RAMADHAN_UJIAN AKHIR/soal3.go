package main

import (
	"fmt"
)

const nProv = 3

// Struktur data untuk menyimpan nama, populasi, dan angka pertumbuhan provinsi
type Provinsi struct {
	Nama        string
	Populasi    int
	Pertumbuhan float64
}

// Fungsi untuk menginput data provinsi
func input_2311102241(provinsi []Provinsi) {
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data provinsi ke-%d:\n", i+1)
		fmt.Print("Nama Provinsi: ")
		fmt.Scan(&provinsi[i].Nama)
		fmt.Print("Jumlah Populasi: ")
		fmt.Scan(&provinsi[i].Populasi)
		fmt.Print("Angka Pertumbuhan (%): ")
		fmt.Scan(&provinsi[i].Pertumbuhan)
	}
}

// Fungsi untuk mencari provinsi dengan angka pertumbuhan tercepat
func provTercepat_2311102241(provinsi []Provinsi) string {
	indexTercepat := 0
	for i := 1; i < nProv; i++ {
		if provinsi[i].Pertumbuhan > provinsi[indexTercepat].Pertumbuhan {
			indexTercepat = i
		}
	}
	return provinsi[indexTercepat].Nama
}

// Fungsi untuk mencari indeks provinsi berdasarkan nama
func IndeksProv_2311102241(provinsi []Provinsi, nama string) int {
	for i := 0; i < nProv; i++ {
		if provinsi[i].Nama == nama {
			return i
		}
	}
	return -1 // Jika provinsi tidak ditemukan
}

// Fungsi untuk menampilkan prediksi populasi berdasarkan pertumbuhan > 2%
func populasi_2311102241(provinsi []Provinsi) {
	fmt.Println("Prediksi populasi pertumbuhan sebanyak 2%:")
	for i := 0; i < nProv; i++ {
		if provinsi[i].Pertumbuhan > 2 {
			prediksi := float64(provinsi[i].Populasi) * (1 + provinsi[i].Pertumbuhan/100)
			fmt.Printf("%s: Populasi saat ini = %d, Prediksi populasi = %.2f\n", provinsi[i].Nama, provinsi[i].Populasi, prediksi)
		}
	}
}

func main() {
	var provinsi [nProv]Provinsi
	var namaCari string

	// Input data
	fmt.Println("Masukkan data untuk 34 provinsi:")
	input_2311102241(provinsi[:])

	// Provinsi dengan pertumbuhan tercepat
	namaTercepat := provTercepat_2311102241(provinsi[:])
	fmt.Printf("Provinsi dengan angka pertumbuhan tercepat: %s\n", namaTercepat)

	// Cari provinsi berdasarkan nama
	fmt.Print("Masukkan nama provinsi yang ingin dicari: ")
	fmt.Scan(&namaCari)

	fmt.Print("\nHasil Pencarian Provinsi\n\n")
	index := IndeksProv_2311102241(provinsi[:], namaCari)
	if index == -1 {
		fmt.Printf("Provinsi dengan nama '%s' tidak ditemukan.\n", namaCari)
	} else {
		fmt.Printf("Provinsi %s memiliki populasi %d dan pertumbuhan %.2f%%\n", provinsi[index].Nama, provinsi[index].Populasi, provinsi[index].Pertumbuhan)
	}

	// Prediksi populasi
	populasi_2311102241(provinsi[:])
}
