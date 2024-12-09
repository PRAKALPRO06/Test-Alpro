//Mahasiswa adalah struct yang berisi NIM, nama, dan array nilai (3 nilai).
//ArrayMahasiswa adalah array statis dengan kapasitas maksimum nMax (51 mahasiswa).

package main

import (
	"fmt"
)

const nMax = 51

type Mahasiswa struct {
	NIM   string
	Nama  string
	Nilai [3]int
}

type ArrayMahasiswa [nMax]Mahasiswa

func main() {
	var dataMahasiswa ArrayMahasiswa
	var n int

	//Input data mahasiswa
	fmt.Print("Masukkan jumlah mahasiswa (maks 51): ")
	fmt.Scan(&n)

	if n > nMax {
		fmt.Println("Jumlah mahasiswa melebihi kapasitas!")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("\nMahasiswa ke-%d:\n", i+1)
		fmt.Print("Masukkan NIM: ")
		fmt.Scan(&dataMahasiswa[i].NIM)
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&dataMahasiswa[i].Nama)
		fmt.Println("Masukkan 3 nilai mahasiswa:")
		for j := 0; j < 3; j++ {
			fmt.Printf("Nilai %d: ", j+1)
			fmt.Scan(&dataMahasiswa[i].Nilai[j])
		}
	}

	//mencari nilai pertama berdasarkan NIM
	var cariNIM string
	fmt.Print("\nMasukkan NIM untuk mencari nilai pertama: ")
	fmt.Scan(&cariNIM)

	found := false
	for i := 0; i < n; i++ {
		if dataMahasiswa[i].NIM == cariNIM {
			fmt.Printf("Nilai pertama mahasiswa dengan NIM %s adalah %d\n", cariNIM, dataMahasiswa[i].Nilai[0])
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
	}

	//Mencari nilai terbesar berdasarkan NIM
	fmt.Print("\nMasukkan NIM untuk mencari nilai terbesar: ")
	fmt.Scan(&cariNIM)

	found = false
	for i := 0; i < n; i++ {
		if dataMahasiswa[i].NIM == cariNIM {
			maks := dataMahasiswa[i].Nilai[0]
			for _, nilai := range dataMahasiswa[i].Nilai {
				if nilai > maks {
					maks = nilai
				}
			}
			fmt.Printf("Nilai terbesar mahasiswa dengan NIM %s adalah %d\n", cariNIM, maks)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
	}

	//Tampilkan hasil pencarian
	fmt.Println("\nData Mahasiswa:")
	for i := 0; i < n; i++ {
		fmt.Printf("NIM: %s, Nama: %s, Nilai: %d, %d, %d\n",
			dataMahasiswa[i].NIM,
			dataMahasiswa[i].Nama,
			dataMahasiswa[i].Nilai[0],
			dataMahasiswa[i].Nilai[1],
			dataMahasiswa[i].Nilai[2],
		)
	}
}
