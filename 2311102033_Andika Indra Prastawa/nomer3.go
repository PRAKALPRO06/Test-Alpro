package main

import (
	"fmt"
	"strings"
)

// Konstanta jumlah provinsi
const nProv = 34

// Tipe data untuk menyimpan data provinsi
type (
	NamaProv   [nProv]string
	PopProv    [nProv]int
	TumbuhProv [nProv]float64
)

// Fungsi untuk menginput data provinsi
func InputData(nama *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan nama provinsi ke-%d: ", i+1)
		fmt.Scanln(&nama[i])
		fmt.Printf("Masukkan populasi provinsi %s: ", nama[i])
		fmt.Scanln(&pop[i])
		fmt.Printf("Masukkan angka pertumbuhan penduduk provinsi %s (dalam desimal): ", nama[i])
		fmt.Scanln(&tumbuh[i])
	}
}

// Fungsi untuk mencari indeks provinsi dengan pertumbuhan tercepat
func ProvinsiTercepat(tumbuh TumbuhProv) int {
	indeks := 0
	maks := tumbuh[0]
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > maks {
			maks = tumbuh[i]
			indeks = i
		}
	}
	return indeks
}

// Fungsi untuk mencari indeks provinsi berdasarkan nama
func IndeksProvinsi(nama NamaProv, target string) int {
	for i := 0; i < nProv; i++ {
		if strings.EqualFold(nama[i], target) {
			return i
		}
	}
	return -1
}

// Prosedur untuk menampilkan prediksi populasi provinsi dengan pertumbuhan > 2%
func Prediksi(nama NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\nPrediksi provinsi dengan pertumbuhan > 2%:")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := float64(pop[i]) * (1 + tumbuh[i])
			fmt.Printf("%s: Populasi tahun depan = %.0f\n", nama[i], prediksi)
		}
	}
}

func main() {
	var (
		nama   NamaProv
		pop    PopProv
		tumbuh TumbuhProv
		cari   string
	)

	// Input data
	fmt.Println("Masukkan data provinsi:")
	InputData(&nama, &pop, &tumbuh)

	// Provinsi dengan pertumbuhan tercepat
	tercepat := ProvinsiTercepat(tumbuh)
	fmt.Printf("\nProvinsi dengan pertumbuhan tercepat: %s\n", nama[tercepat])

	// Cari provinsi berdasarkan nama
	fmt.Println("\nMasukkan nama provinsi yang ingin dicari:")
	fmt.Scanln(&cari)
	indeks := IndeksProvinsi(nama, cari)
	if indeks != -1 {
		fmt.Printf("Provinsi %s ditemukan dengan populasi %d dan pertumbuhan %.2f%%\n",
			nama[indeks], pop[indeks], tumbuh[indeks]*100)
	} else {
		fmt.Printf("Provinsi %s tidak ditemukan\n", cari)
	}

	// Tampilkan prediksi populasi
	Prediksi(nama, pop, tumbuh)
}
