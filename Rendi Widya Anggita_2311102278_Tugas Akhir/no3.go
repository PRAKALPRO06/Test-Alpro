package main

import (
	"fmt"
)

const nProv_2311102278 = 3

type (
	NamaProv   [nProv_2311102278]string
	PopProv    [nProv_2311102278]int
	TumbuhProv [nProv_2311102278]float64
)

// Fungsi untuk menyimpan nama provinsi, jumlah populasi, dan angka pertumbuhan
func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("Masukkan data untuk 3 provinsi:")
	for i := 0; i < nProv_2311102278; i++ {
		fmt.Printf("\nProvinsi ke-%d:\n", i+1)

		// Input nama provinsi
		fmt.Print("Nama Provinsi: ")
		fmt.Scan(&(*prov)[i])

		// Input populasi
		fmt.Print("Jumlah Populasi: ")
		fmt.Scan(&(*pop)[i])

		// Input pertumbuhan
		fmt.Print("Angka Pertumbuhan (%): ")
		fmt.Scan(&(*tumbuh)[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	maxIndex := 0
	maxGrowth := tumbuh[0]

	for i := 1; i < nProv_2311102278; i++ {
		if tumbuh[i] > maxGrowth {
			maxGrowth = tumbuh[i]
			maxIndex = i
		}
	}

	return maxIndex
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("\nProvinsi dengan pertumbuhan di atas 2%:")

	for i := 0; i < nProv_2311102278; i++ {
		if tumbuh[i] > 2.0 {
			prediksiPop := int(float64(pop[i]) * (1 + tumbuh[i]/100))
			fmt.Printf("Nama Provinsi: %s, Prediksi Populasi Tahun Depan: %d\n", prov[i], prediksiPop)
		}
	}
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv_2311102278; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv

	// Input data oleh user
	InputData(&prov, &pop, &tumbuh)

	// Cari provinsi dengan pertumbuhan tercepat
	tercepatIndex := ProvinsiTercepat(tumbuh)
	fmt.Printf("\nProvinsi dengan pertumbuhan tercepat: %s\n", prov[tercepatIndex])

	// Cari indeks provinsi tertentu
	var namaDicari string
	fmt.Print("\nMasukkan nama provinsi yang ingin dicari: ")
	fmt.Scan(&namaDicari)

	index := IndeksProvinsi(prov, namaDicari)
	if index != -1 {
		fmt.Printf("Indeks provinsi %s: %d\n", namaDicari, index)
	} else {
		fmt.Printf("Provinsi %s tidak ditemukan.\n", namaDicari)
	}

	// Prediksi populasi untuk pertumbuhan di atas 2%
	Prediksi(prov, pop, tumbuh)
}
