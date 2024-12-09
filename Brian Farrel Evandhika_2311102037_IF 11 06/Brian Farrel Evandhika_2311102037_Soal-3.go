// 2311102037 Brian Farrel Evandhika IF 11 06
package main

import (
	"fmt"
)

const nProv = 34

type NamaProv_2311102037 [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv_2311102037, pop *PopProv, tumbuh *TumbuhProv) { // fungsi input data mengisi array prov, pop, dan tumbuh dengan data yang diberikan oleh pengguna
	for i := 0; i < nProv; i++ {
		fmt.Printf("masukkan nama provinsi ke-%d: ", i+1)
		fmt.Scanln(&prov[i])

		fmt.Printf("masukkan populasi provinsi %s: ", prov[i])
		fmt.Scanln(&pop[i])

		fmt.Printf("masukkan angka pertumbuhan untuk provinsi %s (dalam persen): ", prov[i])
		fmt.Scanln(&tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int { // fungsi provinsitercepat mengembalikan indeks array tumbuh dengan pertumbuhan penduduk tercepat
	tercepat := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[tercepat] {
			tercepat = i
		}
	}
	return tercepat
}

func Prediksi(prov NamaProv_2311102037, pop PopProv, tumbuh TumbuhProv) {
	// fungsi prediksi menampilkan nama provinsi dan prediksi jumlah penduduknya di tahun depan dengan pertumbuhan di atas 2%
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 2.0 {
			prediksiPop := int(float64(pop[i]) * (1 + tumbuh[i]/100))
			fmt.Printf("%s: %d\n", prov[i], prediksiPop)
		}
	}
}

// IndeksProvinsi mengembalikan indeks array prov untuk provinsi dengan nama tertentu
func IndeksProvinsi(prov NamaProv_2311102037, nama string) int { // fungsi indeksprovinsi mengembalikan indeks array prov untuk provinsi dengan nama tertentu
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func main() { //fungsi utama untuk memanggil fungsi lainya
	var prov NamaProv_2311102037
	var pop PopProv
	var tumbuh TumbuhProv

	InputData(&prov, &pop, &tumbuh)

	tercepatIndex := ProvinsiTercepat(tumbuh)
	fmt.Println("provinsi dengan pertumbuhan tercepat:", prov[tercepatIndex])

	fmt.Print("masukkan nama provinsi yang akan dicari: ")
	var namaProvinsi string
	fmt.Scanln(&namaProvinsi)

	indeks := IndeksProvinsi(prov, namaProvinsi)
	fmt.Printf("indeks dari %s: %d\n", namaProvinsi, indeks)

	fmt.Println("prediksi jumlah penduduk tahun depan untuk provinsi dengan pertumbuhan di atas 2%:")
	Prediksi(prov, pop, tumbuh)
}
