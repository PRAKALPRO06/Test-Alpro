package main

import (
	"fmt"
	"strings"
)

// Konstanta untuk jumlah provinsi
const nProv_2311102166 = 34

func main() {
	// Deklarasi variabel
	var NamaProv [nProv_2311102166]string
	var PopProv [nProv_2311102166]int
	var TumbuhProv [nProv_2311102166]float64
	var ProvinsiDicari string

	// Input data
	InputData(&NamaProv, &PopProv, &TumbuhProv)

	// Input nama provinsi yang akan dicari
	fmt.Println("Masukkan nama provinsi yang akan dicari:")
	fmt.Scanln(&ProvinsiDicari)

	// Temukan provinsi dengan pertumbuhan tercepat
	provinsiTercepat := ProvinsiTercepat(NamaProv, TumbuhProv)
	fmt.Println("Provinsi dengan pertumbuhan tercepat adalah:", provinsiTercepat)

	// Temukan indeks provinsi yang dicari
	index := IndeksProvinsi(NamaProv, ProvinsiDicari)
	if index != -1 {
		fmt.Printf("Indeks provinsi %s adalah: %d\n", ProvinsiDicari, index)
	} else {
		fmt.Printf("Provinsi %s tidak ditemukan.\n", ProvinsiDicari)
	}

	// Prediksi jumlah penduduk untuk provinsi dengan pertumbuhan di atas 2%
	fmt.Println("Prediksi jumlah penduduk untuk provinsi dengan pertumbuhan di atas 2%:")
	Prediksi(NamaProv, PopProv, TumbuhProv)
}

// Subprogram untuk input data dari user
func InputData(NamaProv *[nProv_2311102166]string, PopProv *[nProv_2311102166]int, TumbuhProv *[nProv_2311102166]float64) {
	fmt.Println("Masukkan data provinsi (Nama, Populasi, Pertumbuhan):")
	for i := 0; i < nProv_2311102166; i++ {
		fmt.Printf("Data provinsi ke-%d:\n", i+1)
		fmt.Scan(&NamaProv[i], &PopProv[i], &TumbuhProv[i])
	}
}

// Subprogram untuk menemukan provinsi dengan pertumbuhan penduduk tercepat
func ProvinsiTercepat(NamaProv [nProv_2311102166]string, TumbuhProv [nProv_2311102166]float64) string {
	maxGrowth := TumbuhProv[0]
	index := 0

	for i := 1; i < nProv_2311102166; i++ {
		if TumbuhProv[i] > maxGrowth {
			maxGrowth = TumbuhProv[i]
			index = i
		}
	}

	return NamaProv[index]
}

// Subprogram untuk menemukan indeks provinsi berdasarkan nama
func IndeksProvinsi(NamaProv [nProv_2311102166]string, nama string) int {
	for i := 0; i < nProv_2311102166; i++ {
		if strings.EqualFold(NamaProv[i], nama) {
			return i
		}
	}
	return -1 // Jika tidak ditemukan
}

// Subprogram untuk memprediksi jumlah penduduk berdasarkan pertumbuhan
func Prediksi(NamaProv [nProv_2311102166]string, PopProv [nProv_2311102166]int, TumbuhProv [nProv_2311102166]float64) {
	for i := 0; i < nProv_2311102166; i++ {
		if TumbuhProv[i] > 2.0 {
			// Hitung prediksi populasi
			prediksi := float64(PopProv[i]) * (1 + TumbuhProv[i]/100)
			fmt.Printf("Provinsi %s: %.0f\n", NamaProv[i], prediksi)
		}
	}
}
