package main

import "fmt"

type DataProvinsi struct {
	Nama_2311102264     string
	Populasi_2311102264 int
	Pertumbuhan         float64
}

func main() {
	var data [2]DataProvinsi
	var namaCari string

	// Input data provinsi
	for i := 0; i < 34; i++ {
		fmt.Scanln(&data[i].Nama_2311102264, &data[i].Populasi_2311102264, &data[i].Pertumbuhan)
	}

	// Cari provinsi dengan pertumbuhan tercepat
	indeksTercepat := 0
	for i := 1; i < 34; i++ {
		if data[i].Pertumbuhan > data[indeksTercepat].Pertumbuhan {
			indeksTercepat = i
		}
	}
	fmt.Println(data[indeksTercepat].Nama_2311102264)

	// Cari indeks provinsi berdasarkan nama
	fmt.Scanln(&namaCari)
	indeksCari := -1
	for i := 0; i < 34; i++ {
		if data[i].Nama_2311102264 == namaCari {
			indeksCari = i
			break
		}
	}
	fmt.Println(indeksCari)

	// Prediksi jumlah penduduk tahun depan
	fmt.Println("Prediksi Jumlah Penduduk Tahun Depan (Pertumbuhan > 2%):")
	for _, prov := range data {
		if prov.Pertumbuhan > 0.02 {
			prediksi := float64(prov.Populasi_2311102264) + (float64(prov.Populasi_2311102264) * prov.Pertumbuhan)
			fmt.Printf("- %s: %.0f\n", prov.Nama_2311102264, prediksi)
		}
	}
}
