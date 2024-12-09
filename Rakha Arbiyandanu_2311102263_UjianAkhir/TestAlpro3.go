

package main

import "fmt"

const nProv = 35//digunakan utuk mebatasi data 

type Provinsi struct {//untutk mengatur tipe data dan mengatur untuk fungsinya
	nama       string
	populasi   int
	pertumbuhan float64
}

type arrayProvinsi [nProv]Provinsi


func input_2311102263(n int) arrayProvinsi {//inputan untuk provinsi
	var data arrayProvinsi
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data provinsi ke-%d (Nama Populasi Pertumbuhan): ", i+1)
		fmt.Scan(&data[i].nama, &data[i].populasi, &data[i].pertumbuhan)
	}
	return data
}

func ProvinsiTertinggi(data arrayProvinsi, n int) int {//mencari provinsi dengan pertumabahan tertinggi
	maxIndex := 0
	for i := 1; i < n; i++ {
		if data[i].pertumbuhan > data[maxIndex].pertumbuhan {
			maxIndex = i
		}
	}
	return maxIndex
}

func IndeksProvinsi(data arrayProvinsi, n int, nama string) int {//mencari provinsi berdasarkan nama
	for i := 0; i < n; i++ {
		if data[i].nama == nama {
			return i
		}
	}
	return -1
}


func HitungPopulasi(data arrayProvinsi, indeks int) int {//fungsi untuk menghitung populasi tahun berikutnya
	return int(float64(data[indeks].populasi) * (1 + data[indeks].pertumbuhan/100))
}

func main() {//fungsi untuk menjalankan seluruh program
	var n int
	fmt.Print("Masukkan jumlah provinsi: ")
	fmt.Scan(&n)
	if n > nProv {
		fmt.Printf("Jumlah provinsi tidak boleh lebih dari %d\n", nProv)
		return
	}

	data := input_2311102263(n)

	// Menampilkan provinsi dengan pertumbuhan tertinggi
	tinggiIndex := ProvinsiTertinggi(data, n)
	fmt.Printf("Provinsi dengan pertumbuhan tertinggi adalah %s\n", data[tinggiIndex].nama)

	// Mencari provinsi berdasarkan nama
	var namaCari string
	fmt.Print("Masukkan nama provinsi yang ingin dicari: ")
	fmt.Scan(&namaCari)
	indeksCari := IndeksProvinsi(data, n, namaCari)

	if indeksCari != -1 {
		fmt.Printf("Provinsi ditemukan: %s dengan populasi %d dan pertumbuhan %.2f%%\n", data[indeksCari].nama, data[indeksCari].populasi, data[indeksCari].pertumbuhan)

		// Menghitung populasi di tahun berikutnya
		populasiBaru := HitungPopulasi(data, indeksCari)
		fmt.Printf("Populasi provinsi %s di tahun depan: %d\n", data[indeksCari].nama, populasiBaru)

		// Menampilkan pesan jika pertumbuhan di atas 2%
		if data[indeksCari].pertumbuhan > 2 {
			fmt.Printf("Provinsi %s memiliki pertumbuhan di atas 2%%.\n", data[indeksCari].nama)
		}
	} else {
		fmt.Printf("Provinsi dengan nama %s tidak ditemukan.\n", namaCari)
	}
}
