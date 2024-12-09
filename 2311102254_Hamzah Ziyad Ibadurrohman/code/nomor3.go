package main

import (
	"fmt"
	"strings"
)

const nProv = 34

// Tipe data untuk menyimpan nama provinsi, populasi, dan pertumbuhan
type (
	NamaProv   [nProv]string
	PopProv    [nProv]int
	TumbuhProv [nProv]float64
)

// Procedure untuk membaca data input
func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	fmt.Println("Masukkan data provinsi (nama, populasi, pertumbuhan):")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Data provinsi ke-%d:\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&(*prov)[i])
		fmt.Print("Populasi: ")
		fmt.Scan(&(*pop)[i])
		fmt.Print("Pertumbuhan (%): ")
		fmt.Scan(&(*tumbuh)[i])
	}
}

// Fungsi untuk mencari indeks provinsi dengan pertumbuhan penduduk tercepat
func ProvinsiTercepat(tumbuh TumbuhProv) int {
	maxIndex := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

// Procedure untuk mencetak prediksi jumlah penduduk tahun depan
func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("Prediksi jumlah penduduk dengan pertumbuhan di atas 2%:")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 2.0 {
			prediksi := float64(pop[i]) * (1 + tumbuh[i]/100)
			fmt.Printf("%s: %.0f\n", prov[i], prediksi)
		}
	}
}

// Fungsi untuk mencari indeks provinsi berdasarkan nama
func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if strings.EqualFold(prov[i], nama) {
			return i
		}
	}
	return -1
}

func main() {
	var (
		prov   NamaProv
		pop    PopProv
		tumbuh TumbuhProv
		cari   string
	)

	InputData(&prov, &pop, &tumbuh)

	fmt.Print("Masukkan nama provinsi yang dicari: ")
	fmt.Scan(&cari)

	tercepatIndex := ProvinsiTercepat(tumbuh)
	fmt.Printf("Provinsi dengan pertumbuhan tercepat: %s\n", prov[tercepatIndex])

	indexProvinsi := IndeksProvinsi(prov, cari)
	if indexProvinsi != -1 {
		fmt.Printf("Indeks provinsi %s: %d\n", cari, indexProvinsi)
	} else {
		fmt.Printf("Provinsi %s tidak ditemukan.\n", cari)
	}

	// Menampilkan prediksi jumlah penduduk
	Prediksi(prov, pop, tumbuh)
}
