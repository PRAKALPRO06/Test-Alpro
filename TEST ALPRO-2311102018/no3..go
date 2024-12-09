package main

import (
	"fmt"
	"strings"
)

// Maksimum jumlah provinsi
const nProv = 34

// Tipe data
type (
	NamaProv   [nProv]string
	PopProv    [nProv]int
	TumbuhProv [nProv]float64
)

// Fungsi untuk menginput data provinsi
func InputData_2311102018(nama *NamaProv, pop *PopProv, tumbuh *TumbuhProv, jumlah *int) {
	fmt.Printf("Berapa jumlah provinsi yang ingin diinput? (max %d): ", nProv)
	fmt.Scanln(jumlah)
	if *jumlah > nProv || *jumlah <= 0 {
		fmt.Println("Jumlah provinsi tidak valid. Gunakan nilai antara 1 hingga", nProv)
		return
	}

	for i := 0; i < *jumlah; i++ {
		fmt.Printf("Masukkan nama provinsi ke-%d: ", i+1)
		fmt.Scanln(&nama[i])
		fmt.Printf("Masukkan populasi provinsi %s: ", nama[i])
		fmt.Scanln(&pop[i])
		fmt.Printf("Masukkan angka pertumbuhan penduduk provinsi %s: ", nama[i])
		fmt.Scanln(&tumbuh[i])
	}
}

// Fungsi untuk mencari provinsi dengan pertumbuhan tercepat
func ProvinsiTercepat(tumbuh TumbuhProv, jumlah int) int {
	indeks := 0
	maks := tumbuh[0]
	for i := 1; i < jumlah; i++ {
		if tumbuh[i] > maks {
			maks = tumbuh[i]
			indeks = i
		}
	}
	return indeks
}

// Fungsi untuk mencari indeks provinsi berdasarkan nama
func IndeksProvinsi(nama NamaProv, target string, jumlah int) int {
	for i := 0; i < jumlah; i++ {
		if strings.EqualFold(nama[i], target) {
			return i
		}
	}
	return -1
}

// Prosedur untuk menampilkan prediksi populasi provinsi
func Prediksi(nama NamaProv, pop PopProv, tumbuh TumbuhProv, jumlah int) {
	fmt.Println("Prediksi provinsi dengan pertumbuhan > 2%:")
	for i := 0; i < jumlah; i++ {
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
		jumlah int
		cari   string
	)

	// Input data
	InputData_2311102018(&nama, &pop, &tumbuh, &jumlah)
	if jumlah <= 0 {
		fmt.Println("Tidak ada data yang diinput. Program selesai.")
		return
	}

	// Provinsi dengan pertumbuhan tercepat
	tercepat := ProvinsiTercepat(tumbuh, jumlah)
	fmt.Printf("Provinsi dengan pertumbuhan tercepat: %s\n", nama[tercepat])

	// Cari provinsi berdasarkan nama
	fmt.Println("Masukkan nama provinsi yang ingin dicari:")
	fmt.Scanln(&cari)
	indeks := IndeksProvinsi(nama, cari, jumlah)
	if indeks != -1 {
		fmt.Printf("Provinsi %s ditemukan pada indeks %d\n", cari, indeks)
	} else {
		fmt.Printf("Provinsi %s tidak ditemukan\n", cari)
	}

	// Tampilkan prediksi populasi
	Prediksi(nama, pop, tumbuh, jumlah)
}