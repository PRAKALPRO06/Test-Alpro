// 2311102037 Brian Farrel Evandhika IF 11 06
package main

import (
	"fmt"
)

const nMax int = 51

type mahasiswa struct {
	NIM_2311102037 string
	nama           string
	nilai          int
}

type arrayMahasiswa [nMax]mahasiswa

func main() { //fungsi utama untuk memanggil fungsi lainya
	var N int
	var data arrayMahasiswa

	fmt.Print("masukan jumlah data mahasiswa: ") //memasukan data mahasiswa
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Printf("masukan NIM, nama, dan nilai mahasiswa ke-%d: ", i+1)
		fmt.Scan(&data[i].NIM_2311102037, &data[i].nama, &data[i].nilai)
	}

	var cariNIM_2311102037 string // mencari nim mahasiswa yang ingin dicari
	fmt.Print("masukkan NIM yang ingin dicari:")
	fmt.Scan(&cariNIM_2311102037)

	nilaiPertama := cariNilaiPertama(data, N, cariNIM_2311102037)
	nilaiTerbesar := cariNilaiTerbesar(data, N, cariNIM_2311102037)

	if nilaiPertama != -1 {
		fmt.Printf("nilai pertama mahasiswa dengan NIM %s adalah %d\n", cariNIM_2311102037, nilaiPertama)
	} else {
		fmt.Printf("mahasiswa dengan NIM %s tidak ditemukan\n", cariNIM_2311102037)
	}

	if nilaiTerbesar != -1 {
		fmt.Printf("nilai terbesar mahasiswa dengan NIM %s adalah %d\n", cariNIM_2311102037, nilaiTerbesar)
	} else {
		fmt.Printf("mahasiswa dengan NIM %s tidak ditemukan\n", cariNIM_2311102037)
	}
}

func cariNilaiPertama(data arrayMahasiswa, N int, NIM_2311102037 string) int { //fungsi untuk mencari nilai pertama pada nim mahasiswa
	for i := 0; i < N; i++ {
		if data[i].NIM_2311102037 == NIM_2311102037 {
			return data[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(data arrayMahasiswa, N int, NIM_2311102037 string) int { //fungsi untuk mencari nilai terbesar pada nim mahasiswa
	var nilaiTerbesar int = -1
	for i := 0; i < N; i++ {
		if data[i].NIM_2311102037 == NIM_2311102037 {
			nilaiTerbesar = data[i].nilai
		}
	}
	return nilaiTerbesar
}
