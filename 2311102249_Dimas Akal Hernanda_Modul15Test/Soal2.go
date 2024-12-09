package main

import (
	"fmt"
)

type Mahasiswa struct {
	NIM   string
	Nama  string
	Nilai int
}

const nMax = 51

type ArrayMahasiswa [nMax]Mahasiswa

func masukkanData(n int, data *ArrayMahasiswa) {
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data mahasiswa ke-%d:\n", i+1)
		fmt.Print("NIM: ")
		fmt.Scan(&data[i].NIM)
		fmt.Print("Nama: ")
		fmt.Scan(&data[i].Nama)
		fmt.Print("Nilai: ")
		fmt.Scan(&data[i].Nilai)
	}
}

func cariNilaiPertama(data ArrayMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if data[i].NIM == nim {
			return data[i].Nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(data ArrayMahasiswa, n int, nim string) int {
	max := -1
	for i := 0; i < n; i++ {
		if data[i].NIM == nim && data[i].Nilai > max {
			max = data[i].Nilai
		}
	}
	return max
}

func main() {
	var n int
	var nim string
	var data ArrayMahasiswa

	fmt.Print("Masukkan jumlah mahasiswa_2311102249: ")
	fmt.Scan(&n)

	if n > nMax {
		fmt.Printf("Jumlah mahasiswa melebihi kapasitas maksimum (%d).\n", nMax)
		return
	}

	masukkanData(n, &data)

	fmt.Print("Masukkan NIM yang ingin dicari: ")
	fmt.Scan(&nim)

	nilaiPertama := cariNilaiPertama(data, n, nim)
	if nilaiPertama != -1 {
		fmt.Printf("Nilai pertama mahasiswa dengan NIM %s adalah %d.\n", nim, nilaiPertama)
	} else {
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan.\n", nim)
	}

	nilaiTerbesar := cariNilaiTerbesar(data, n, nim)
	if nilaiTerbesar != -1 {
		fmt.Printf("Nilai terbesar mahasiswa dengan NIM %s adalah %d.\n", nim, nilaiTerbesar)
	} else {
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan.\n", nim)
	}
}
