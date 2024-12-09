package main

import (
	"fmt"
	"strings"
)

const nProv = 34

// Struck untuk menyimpan Informasi Provinsi
type provinsi_2311102327 struct {
	nama       string
	populasi   int
	pertumbuhan float64
}

// Fungsi untuk menangai proses input dari pengguna
func inputData() []provinsi_2311102327 {
	data := make([]provinsi_2311102327, nProv)
	for i := 0; i < nProv; i++ {
		fmt.Printf("Provinsi ke-%d\n", i+1)
		fmt.Print("Masukkan nama provinsi: ")
		fmt.Scan(&data[i].nama)

		fmt.Print("Masukkan populasi (dalam jutaan): ")
		fmt.Scan(&data[i].populasi)

		fmt.Print("Masukkan pertumbuhan (dalam persen): ")
		fmt.Scan(&data[i].pertumbuhan)
		fmt.Println()
	}
	return data
}

// Fungsi perulangan untuk indeks lalu jika sesuai, maka akan dikembalikan
func provinsiTercepat(data []provinsi_2311102327) int {
	tercepat := 0
	for i := 1; i < len(data); i++ {
		if data[i].pertumbuhan > data[tercepat].pertumbuhan {
			tercepat = i
		}
	}
	return tercepat
}

// Fungsi untuk mencari indeks berdasarkan nama
func indeksProvinsi(data []provinsi_2311102327, nama string) int {
	for i, prov := range data {
		if strings.EqualFold(prov.nama, nama) {
			return i
		}
	}
	return -1
}

// Fungsi untuk menampilkan prediksi populasi jika pertumbuhan di atas 2%
func prediksiPopulasi(data []provinsi_2311102327) {
	fmt.Println("\nPrediksi populasi untuk provinsi dengan pertumbuhan >= 2%:")
	for _, prov := range data {
		if prov.pertumbuhan >= 2 {
			prediksi := float64(prov.populasi) * (1 + prov.pertumbuhan/100)
			fmt.Printf("Provinsi: %s, Populasi saat ini: %d juta, Prediksi tahun depan: %.2f juta\n",
				prov.nama, prov.populasi, prediksi)
		}
	}
}

func main() {
	fmt.Println("(Data provinsi)")
	data := inputData()

	tercepat := provinsiTercepat(data)
	fmt.Printf("\nProvinsi dengan pertumbuhan tercepat adalah: %s (%.2f%%)\n",
		data[tercepat].nama, data[tercepat].pertumbuhan)

	var namaCari string
	fmt.Print("\nMasukkan nama provinsi untuk dicari: ")
	fmt.Scan(&namaCari)
	index := indeksProvinsi(data, namaCari)
	if index != -1 {
		fmt.Printf("Data provinsi ditemukan: %s, Populasi: %d juta, Pertumbuhan: %.2f%%\n",
			data[index].nama, data[index].populasi, data[index].pertumbuhan)
	} else {
		fmt.Println("Provinsi tidak ditemukan.")
	}

	prediksiPopulasi(data)
}