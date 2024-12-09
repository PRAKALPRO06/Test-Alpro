package main

import "fmt"

// Fungsi untuk memeriksa apakah elemen sudah ada di himpunan
func simpan(set_2311102278 []int, val int) bool {
	for _, v := range set_2311102278 {
		if v == val {
			return true
		}
	}
	return false
}

// Fungsi untuk membaca masukan dan membentuk himpunan
// Program akan berhenti jika angka yang diinput sudah ada di himpunan
func masukkan() []int {
	var set_2311102278 []int
	var n int
	fmt.Println("Masukkan angka (akhiri dengan -1):")
	for {
		fmt.Scan(&n)
		if n == -1 { // Kondisi berhenti
			break
		}
		if simpan(set_2311102278, n) {
			fmt.Println("Angka sudah diinput sebelumnya. Program dihentikan.")
			break
		}
		set_2311102278 = append(set_2311102278, n)
	}
	return set_2311102278
}

// Fungsi untuk menemukan irisan dari dua himpunan
func cari(set1, set2 []int) []int {
	var isi []int
	for _, v := range set1 {
		if simpan(set2, v) {
			isi = append(isi, v)
		}
	}
	return isi
}

func main() {
	fmt.Println("Input himpunan pertama:")
	set1 := masukkan()
	fmt.Println("Input himpunan kedua:")
	set2 := masukkan()

	isi := cari(set1, set2)

	fmt.Println("Irisan dari himpunan pertama dan kedua adalah:")
	fmt.Println(isi)
}
