package main

import "fmt"


// Tipe data set untuk simpan angka-angka
type set [2022]int

// Fungsi cek udah ada belum si angka di dalam array
func exist(T set, n int, val int) bool {
	// Loop cari di dalam array T
	for i := 0; i < n; i++ {
		if T[i] == val { // Kalau ada yang sama, balikin true
			return true
		}
	}
	return false // Kalau nggak ada, balikin false
}

// Fungsi buat masukin angka ke set
func inputSet(T *set, n *int) {
	var input int
	*n = 0 // Mulai dari nol ya!
	for {
		_, err := fmt.Scan(&input) // Baca angka dari user
		if err != nil {            // Kalau user nggak ngasih angka lagi, selesai deh
			break
		}
		if exist(*T, *n, input) { // Kalau angka udah ada, stop
			break
		}
		if *n < len(T) { // Kalau arraynya belum penuh
			T[*n] = input // Simpen angkanya
			*n++          // Tambah jumlah elemen
		} else { // Kalau penuh, stop juga
			break
		}
	}
}

// Fungsi cari irisan (angka yang sama) antara dua set
func findIntersection(T1, T2 set, n, m int, _2311102166 *set, h *int) {
	*h = 0 // Mulaiin si set hasil kosong ya
	for i := 0; i < n; i++ {
		if exist(T2, m, T1[i]) { // Kalau angka T1[i] ada di T2 juga
			_2311102166[*h] = T1[i] // Masukin ke set hasil
			*h++                    // Tambah jumlah elemen hasil
		}
	}
}

// Fungsi buat print isi set
func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ") // Pisah pakai spasi ya
		}
		fmt.Print(T[i]) // Print angkanya
	}
	fmt.Println() // Baris baru setelah selesai print
}

func main() {
	var s1, s2, _2311102166 set // Ada tiga set: s1, s2, dan hasil (_2311102166)
	var n1, n2, n3 int          // Jumlah elemen di masing-masing set

	// Minta user masukin anggota himpunan pertama
	fmt.Println("Masukkan anggota himpunan pertama:")
	inputSet(&s1, &n1)
	// Minta user masukin anggota himpunan kedua
	fmt.Println("Masukkan anggota himpunan kedua:")
	inputSet(&s2, &n2)

	// Cari irisan kedua himpunan
	findIntersection(s1, s2, n1, n2, &_2311102166, &n3)

	// Print hasil irisan
	fmt.Println("Irisan dari kedua himpunan:")
	printSet(_2311102166, n3)
}