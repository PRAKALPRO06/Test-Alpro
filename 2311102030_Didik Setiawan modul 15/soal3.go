package main

import "fmt"

// Definisi struct propinsi dengan atribut nama_propinsi, nim, kelas, jurusan, dan jumlah_penduduk_2311102030
type propinsi struct {
	nama_propinsi, jurusan     string
	jumlah_penduduk_2311102030 float64
}

// Definisi tipe data array propinsi dengan kapasitas maksimal 2023
type arrMhs [2023]propinsi

// Fungsi untuk mencari jumlah_penduduk_2311102030 tertinggi dalam array propinsi
func jumlah_penduduk_2311102030(T arrMhs, n int) float64 {
	var tertinggi float64 = T[0].jumlah_penduduk_2311102030
	var j int = 1
	for j < n {
		if tertinggi < T[j].jumlah_penduduk_2311102030 {
			tertinggi = T[j].jumlah_penduduk_2311102030
		}
		j = j + 1
	}
	return tertinggi
}

// Fungsi main untuk mengisi data propinsi dan mencari jumlah_penduduk_2311102030 tertinggi
func main() {
	var n int
	var dataMhs arrMhs

	// Meminta input jumlah propinsi
	fmt.Print("Masukkan jumlah propinsi (maks 2023): ")
	fmt.Scan(&n)

	// Validasi jumlah propinsi yang dimasukkan
	if n < 1 || n > 2023 {
		fmt.Println("Jumlah propinsi harus antara 1 dan 2023.")
		return
	}

	// Mengisi data propinsi
	for i := 0; i < n; i++ {
		fmt.Printf("\nMasukkan data propinsi ke-%d\n", i+1)
		fmt.Print("Nama_propinsi: ")
		fmt.Scan(&dataMhs[i].nama_propinsi)
		fmt.Print("jumlah_penduduk_2311102030: ")
		fmt.Scan(&dataMhs[i].jumlah_penduduk_2311102030)
	}

	// Mencari dan menampilkan jumlah_penduduk_2311102030 tertinggi
	tertinggi := jumlah_penduduk_2311102030(dataMhs, n)
	fmt.Printf("\njumlah_penduduk_2311102030 tertinggi dari %d propinsi adalah: %.2f\n", n,
		tertinggi)
}
