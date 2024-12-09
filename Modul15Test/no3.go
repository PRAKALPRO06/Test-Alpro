package main

import (
	"fmt"
)

const nProvinsi = 5

type (
	NamaProvinsi   [nProvinsi]string
	PopProvinsi    [nProvinsi]int
	TumbuhProvinsi [nProvinsi]float64
)

// InputData: Membaca data nama provinsi, populasi, dan angka pertumbuhan dari pengguna
func InputData(prov *NamaProvinsi, pop_2311102273 *PopProvinsi, tumbuh *TumbuhProvinsi, namaCari *string) {
	for i := 0; i < nProvinsi; i++ {
		fmt.Printf("Masukkan nama provinsi ke-%d: ", i+1)
		fmt.Scanln(&prov[i])
		fmt.Printf("Masukkan populasi provinsi ke-%d: ", i+1)
		fmt.Scanln(&pop_2311102273[i])
		fmt.Printf("Masukkan angka pertumbuhan provinsi ke-%d: ", i+1)
		fmt.Scanln(&tumbuh[i])
	}
	fmt.Printf("Masukkan nama provinsi yang ingin dicari: ")
	fmt.Scanln(namaCari)
}

// ProvinsiTercepat: Menemukan indeks provinsi dengan angka pertumbuhan tertinggi
func ProvinsiTercepat(tumbuh TumbuhProvinsi) int {
	maxIndex := 0
	for i := 1; i < nProvinsi; i++ {
		if tumbuh[i] > tumbuh[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

// IndeksProvinsi: Mengembalikan indeks provinsi berdasarkan nama yang dicari
func IndeksProvinsi(prov NamaProvinsi, nama string) int {
	for i := 0; i < nProvinsi; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

// Prediksi: Menampilkan prediksi jumlah penduduk untuk provinsi dengan pertumbuhan > 2%
func Prediksi(prov NamaProvinsi, pop_2311102273 PopProvinsi, tumbuh TumbuhProvinsi) {
	fmt.Println("Prediksi jumlah penduduk tahun depan untuk provinsi dengan pertumbuhan di atas 2%:")
	for i := 0; i < nProvinsi; i++ {
		if tumbuh[i] > 2 {
			prediksi := float64(pop_2311102273[i]) * (1 + tumbuh[i]/100)
			fmt.Printf("%s: %.0f\n", prov[i], prediksi)
		}
	}
}

func main() {
	var prov NamaProvinsi
	var pop_2311102273 PopProvinsi
	var tumbuh TumbuhProvinsi
	var namaCari string

	// Memasukkan data provinsi
	InputData(&prov, &pop_2311102273, &tumbuh, &namaCari)

	// Menemukan provinsi dengan angka pertumbuhan tercepat
	tercepat := ProvinsiTercepat(tumbuh)
	fmt.Printf("Provinsi dengan pertumbuhan tercepat: %s\n", prov[tercepat])

	// Mencari indeks provinsi yang diinputkan oleh pengguna
	indeksCari := IndeksProvinsi(prov, namaCari)
	if indeksCari == -1 {
		fmt.Printf("Provinsi dengan nama '%s' tidak ditemukan.\n", namaCari)
	} else {
		fmt.Printf("Indeks provinsi '%s': %d\n", namaCari, indeksCari)
	}

	// Menampilkan prediksi jumlah penduduk
	Prediksi(prov, pop_2311102273, tumbuh)
}
