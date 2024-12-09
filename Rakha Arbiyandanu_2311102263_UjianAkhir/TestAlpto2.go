package main

import "fmt"

const nMax = 51//sebagai batas data mahasiswawa yang dapat di inputkan

type mahasiswa struct {//mengunakan struct untuk mengatur fungsi fuungsi yang ada
	NIM   string
	nama  string
	nilai int
}

// Fungsi untuk membaca data mahasiswa dan menggunakan fungsi slice agar data agar sesuai dengan data yang di input
func inputMahasiswa_2311102263(n int) []mahasiswa {
	data := make([]mahasiswa, n) 
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data mahasiswa ke-%d (NIM nama nilai): ", i+1)
		fmt.Scan(&data[i].NIM, &data[i].nama, &data[i].nilai)
	}
	return data
}

// Fungsi mencari data mahasiswa dengan nim
func cariNilaiPertama(data []mahasiswa, nim string) (int, bool) {
	for _, mhs := range data {
		if mhs.NIM == nim {
			return mhs.nilai, true
		}
	}
	return 0, false
}

// Fungsi untuk mencari nilai terbesaar dalam dengan bantuan nim
func cariNilaiTerbesar(data []mahasiswa, nim string) (int, bool) {
	max := -1
	found := false
	for _, mhs := range data {
		if mhs.NIM == nim {
			found = true
			if mhs.nilai > max {
				max = mhs.nilai
			}
		}
	}
	return max, found
}

func main() {//fungsi unutk menjalankan seluruh program 
	var n int
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&n)

	if n <= 0 || n > nMax {
		fmt.Printf("Jumlah mahasiswa harus antara 1 hingga %d.\n", nMax)
		return
	}

	data := inputMahasiswa_2311102263(n)

	var nim string
	fmt.Print("Masukkan NIM untuk pencarian: ")
	fmt.Scan(&nim)

	// Cari nilai pertama
	if nilaiPertama, found := cariNilaiPertama(data, nim); found {
		fmt.Printf("Nilai pertama mahasiswa dengan NIM %s adalah %d\n", nim, nilaiPertama)
	} else {
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan\n", nim)
	}

	// Cari nilai terbesar
	if nilaiTerbesar, found := cariNilaiTerbesar(data, nim); found {
		fmt.Printf("Nilai terbesar mahasiswa dengan NIM %s adalah %d\n", nim, nilaiTerbesar)
	} else {
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan\n", nim)
	}
}
