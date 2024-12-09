package main

import (
	"fmt"
	"strings"
)

// Konstanta jumlah provinsi
const nProv = 3

// Tipe data
type (
	NamaProv [nProv]string
	PopProv [nProv]int
	TumbuhProv [nProv]float64
)

// Fungsi untuk menginput data provinsi
func InputData(nama_2311102019 *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan nama provinsi ke-%d: ", i+1)
		fmt.Scanln(&nama_2311102019[i])
		fmt.Printf("Masukkan populasi provinsi %s: ", nama_2311102019[i])
		fmt.Scanln(&pop[i])
		fmt.Printf("Masukkan angka pertumbuhan penduduk provinsi %s: ", nama_2311102019[i])
		fmt.Scanln(&tumbuh[i])
	}
}

// Fungsi untuk mencari provinsi dengan pertumbuhan tercepat
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
		if strings.EqualFold(nama_2311102019[i], target) {
			return i
		}
	}
	return -1
}

// Prosedur untuk menampilkan prediksi populasi provinsi
func Prediksi(nama_2311102019 NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("Prediksi provinsi dengan pertumbuhan > 2%:")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := float64(pop[i]) * (1 + tumbuh[i])
			fmt.Printf("%s: Populasi tahun depan = %.0f\n", nama_2311102019[i], prediksi)
		}
	}
}

func main() {
	var (
		nama_2311102019 NamaProv
		pop PopProv
		tumbuh TumbuhProv
		cari   string
	)

	// Input data
	InputData(&nama_2311102019, &pop, &tumbuh)

	// Provinsi dengan pertumbuhan tercepat
	tercepat := ProvinsiTercepat(tumbuh)
	fmt.Printf("Provinsi dengan pertumbuhan tercepat: %s\n", nama_2311102019[tercepat])

	// Cari provinsi berdasarkan nama
	fmt.Println("Masukkan nama provinsi yang ingin dicari: ")
	fmt.Scanln(&cari)
	indeks := IndeksProvinsi(nama_2311102019, cari)
	if indeks != -1 {
		fmt.Printf("Provinsi %s ditemukan pada indeks %d\n", cari, indeks)
	} else {
		fmt.Printf("Provinsi %s tidak ditemukan\n", cari)
	}

	// Tampilkan prediksi populasi
	Prediksi(nama_2311102019, pop, tumbuh)
}