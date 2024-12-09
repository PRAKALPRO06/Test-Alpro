package main

import (
	"fmt"
)

const nProv = 34 // Jumlah maksimum provinsi

type NamaProv233 [nProv]string
type PopulProv [nProv]int
type TumbuhProv [nProv]float64

// Procedure untuk menginput data
func InputData(prov *NamaProv233, pop *PopulProv, tumbuh *TumbuhProv) {
	fmt.Println("Masukkan data untuk 34 provinsi:")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Provinsi ke-%d (Nama Populasi Pertumbuhan): ", i+1)
		fmt.Scan(&(*prov)[i], &(*pop)[i], &(*tumbuh)[i])
	}
}

// Function untuk mencari indeks provinsi dengan pertumbuhan tercepat
func ProvinsiTercepat(tumbuh TumbuhProv) int {
	maxIndex := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

// Procedure untuk menampilkan prediksi populasi tahun depan (pertumbuhan > 2%)
func Prediksi(prov NamaProv233, pop PopulProv, tumbuh TumbuhProv) {
	fmt.Println("Prediksi populasi tahun depan untuk provinsi dengan pertumbuhan di atas 2%:")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := float64(pop[i]) * (1 + tumbuh[i])
			fmt.Printf("%s: %.0f\n", prov[i], prediksi)
		}
	}
}

// Function untuk mencari indeks provinsi berdasarkan nama
func IndeksProvinsi(prov NamaProv233, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1 // Jika nama provinsi tidak ditemukan
}

func main() {
	var prov NamaProv233
	var pop PopulProv
	var tumbuh TumbuhProv
	var cariNama string
	// Input data
	InputData(&prov, &pop, &tumbuh)
	indexTercepat := ProvinsiTercepat(tumbuh)
	fmt.Printf("Provinsi dengan pertumbuhan tercepat: %s\n", prov[indexTercepat])
	fmt.Print("Masukkan nama provinsi yang ingin dicari: ")
	fmt.Scan(&cariNama)
	indexNama := IndeksProvinsi(prov, cariNama)
	if indexNama != -1 {
		fmt.Printf("Indeks provinsi %s: %d\n", cariNama, indexNama)
	} else {
		fmt.Printf("Provinsi %s tidak ditemukan\n", cariNama)
	}
	Prediksi(prov, pop, tumbuh)
}
