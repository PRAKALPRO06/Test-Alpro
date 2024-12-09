package main
import (
	"fmt"
	"strings"
)
const nProv = 34
type (
	NamaProv [nProv]string
	PopProv [nProv]int
	TumbuhProv [nProv]float64
)

// Fungsi untuk menginputkan data provinsi 
func InputData(nama *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan nama provinsi ke-%d: ", i+1)
		fmt.Scanln(&nama[i])
		fmt.Printf("Masukkan populasi provinsi %s: ", nama[i])
		fmt.Scanln(&pop[i])
		fmt.Printf("Masukkan angka pertumbuhan penduduk provinsi %s: ", nama[i])
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

// Fungsi untuk mencari indeks provinsi berdasarkan nama provinsi yang telah diinputkan
func IndeksProvinsi(nama NamaProv, target string) int {
	for i := 0; i < nProv; i++ {
		if strings.EqualFold(nama[i], target) {
			return i
		}
	}
	return -1
}

// fungsi untuk menampilkan prediksi populasi provinsi 
func Prediksi(nama NamaProv, pop PopProv, tumbuh TumbuhProv) {
	fmt.Println("Prediksi provinsi dengan pertumbuhan > 2%:")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksi := float64(pop[i]) * (1 + tumbuh[i])
			fmt.Printf("%s: Populasi tahun depan = %.0f\n", nama[i], prediksi)
		}
	}
}

func main() {
	var (
		nama_2311102011   NamaProv
		pop    PopProv
		tumbuh TumbuhProv
		cari   string
	)
	InputData(&nama_2311102011, &pop, &tumbuh)
	tercepat := ProvinsiTercepat(tumbuh)

	fmt.Printf("Provinsi dengan pertumbuhan tercepat: %s\n", nama_2311102011[tercepat])
	fmt.Println("Masukkan nama provinsi yang ingin dicari:")
	fmt.Scanln(&cari)
	
	indeks := IndeksProvinsi(nama_2311102011, cari)
	if indeks != -1 {
		fmt.Printf("Provinsi %s ditemukan pada indeks %d\n", cari, indeks)
	} else {
		fmt.Printf("Provinsi %s tidak ditemukan\n", cari)
	}
	Prediksi(nama_2311102011, pop,tumbuh)
}