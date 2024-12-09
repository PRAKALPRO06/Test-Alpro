package main

import (
	"fmt"
)

type Mahasiswa struct {
	nim   string
	nama  string
	nilai int
}

const nMax int = 51

type arrMahasiswa [nMax]Mahasiswa

func find_nilai_pertama(nim_231102211 string, arr arrMahasiswa, nMahasiswa int) int {
	for i := 0; i < nMahasiswa; i++ {
		if arr[i].nim == nim_231102211 {
			return arr[i].nilai
		}
	}
	return -1
}
func find_nilai_terbesar(nim string, arr arrMahasiswa, nMahasiswa int) int {
	nilai_max := -1
	for i := 0; i < nMahasiswa; i++ {
		if arr[i].nim == nim {
			if arr[i].nilai > nilai_max {
				nilai_max = arr[i].nilai
			}
		}
	}
	if nilai_max != 1 {
		return nilai_max
	} else {
		return -1
	}
}

func tampilkan_nilai_pertama_dan_terbesar(nim string, arr arrMahasiswa, nMahasiswa int) {
	fmt.Println("Nilai pertama: ", find_nilai_pertama(nim, arr, nMahasiswa))
	fmt.Println("Nilai terbesar: ", find_nilai_terbesar(nim, arr, nMahasiswa))
}

func main() {

	var arr arrMahasiswa
	nMahasiswa := 0

	var n int
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&n)
	fmt.Println("Masukkan nim, nama dan nilai mahasiswa di setiap baris (dipisahkan dengan spasi): ")

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i].nim, &arr[i].nama, &arr[i].nilai)
		nMahasiswa++
	}

	var nim_cari string
	fmt.Print("Masukkan nim mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya: ")
	fmt.Scan(&nim_cari)

	tampilkan_nilai_pertama_dan_terbesar(nim_cari, arr, nMahasiswa)

}
