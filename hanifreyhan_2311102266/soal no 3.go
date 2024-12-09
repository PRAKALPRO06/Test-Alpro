package main

import (
	"fmt"
	"strings"
)

type Provinsi struct {
	Nama         string
	Populasi     int
	Pertumbuhan float64
}

func main() {
	const nProv = 34
	var provinsi [nProv]Provinsi

	// Input data
	fmt.Println("Masukkan data 34 provinsi dalam format: Nama Populasi Pertumbuhan")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Provinsi %d: ", i+1)
		var nama string
		var populasi int
		var pertumbuhan float64
		fmt.Scanf("%s %d %f\n", &nama, &populasi, &pertumbuhan)
		provinsi[i] = Provinsi{Nama: nama, Populasi: populasi, Pertumbuhan: pertumbuhan}
	}

	// Input nama provinsi yang dicari
	fmt.Println("Masukkan nama provinsi yang ingin dicari:")
	var cariNama string
	fmt.Scanln(&cariNama)

	// Cari provinsi dengan pertumbuhan tercepat
	tercepat := CariProvinsiTercepat(provinsi[:])
	fmt.Printf("Provinsi dengan pertumbuhan tercepat: %s\n", tercepat.Nama)

	// Cari indeks provinsi berdasarkan nama
	indeks := CariIndeksProvinsi(provinsi[:], cariNama)
	if indeks != -1 {
		fmt.Printf("Indeks provinsi %s: %d\n", cariNama, indeks+1)
	} else {
		fmt.Printf("Provinsi %s tidak ditemukan.\n", cariNama)
	}

	// Prediksi jumlah penduduk tahun depan untuk provinsi dengan pertumbuhan di atas 2%
	fmt.Println("Prediksi jumlah penduduk tahun depan untuk provinsi dengan pertumbuhan > 2%:")
	PrediksiPenduduk(provinsi[:])
}

func CariProvinsiTercepat(provinsi []Provinsi) Provinsi {
	tercepat := provinsi[0]
	for _, p := range provinsi {
		if p.Pertumbuhan > tercepat.Pertumbuhan {
			tercepat = p
		}
	}
	return tercepat
}

func CariIndeksProvinsi(provinsi []Provinsi, nama string) int {
	for i, p := range provinsi {
		if strings.EqualFold(p.Nama, nama) {
			return i
		}
	}
	return -1
}

func PrediksiPenduduk(provinsi []Provinsi) {
	for _, p := range provinsi {
		if p.Pertumbuhan > 2 {
			prediksi := float64(p.Populasi) * (1 + p.Pertumbuhan/100)
			fmt.Printf("%s: %.0f\n", p.Nama, prediksi)
		}
	}
}