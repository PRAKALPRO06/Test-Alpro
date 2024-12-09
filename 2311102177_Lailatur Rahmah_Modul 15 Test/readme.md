<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL XV</strong></h2>
<h2 align="center"><strong> TEST </strong></h2>

<br>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  LAILATUR RAHMAH / 2311102177<br>
  S1 IF 11 06
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Abednego Dwi Septiadi
  
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

## Soal No 1

#### Source Code

```go
package main

import (
	"fmt"
)

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
func findIntersection(T1, T2 set, n, m int, _2311102177 *set, h *int) {
	*h = 0 // Mulaiin si set hasil kosong ya
	for i := 0; i < n; i++ {
		if exist(T2, m, T1[i]) { // Kalau angka T1[i] ada di T2 juga
			_2311102177[*h] = T1[i] // Masukin ke set hasil
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
	var s1, s2, _2311102177 set // Ada tiga set: s1, s2, dan hasil (_2311102177)
	var n1, n2, n3 int          // Jumlah elemen di masing-masing set

	// Minta user masukin anggota himpunan pertama
	fmt.Println("Masukkan anggota himpunan pertama:")
	inputSet(&s1, &n1)
	// Minta user masukin anggota himpunan kedua
	fmt.Println("Masukkan anggota himpunan kedua:")
	inputSet(&s2, &n2)

	// Cari irisan kedua himpunan
	findIntersection(s1, s2, n1, n2, &_2311102177, &n3)

	// Print hasil irisan
	fmt.Println("Irisan dari kedua himpunan:")
	printSet(_2311102177, n3)
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/5dd238f6-0034-4405-b371-5c294b70191e)


## Soal No 2

#### Source Code

```go
package main

import (
	"fmt"
)

// Nih batas maksimal mahasiswa, gak boleh lebih dari 51 ya!
const nMax = 51

// Struktur buat nyimpen data mahasiswa
type mahasiswa struct {
	NIM   string // NIM anak mahasiswa
	nama  string // Nama anak mahasiswa
	nilai int    // Nilai yang didapet
}

// Ini array buat nampung data semua mahasiswa
type arrayMahasiswa [nMax]mahasiswa

// Fungsi buat masukin data mahasiswa ke array
func inputMahasiswa(data *arrayMahasiswa, _2311102177 *int) {
	fmt.Print("Masukkan jumlah mahasiswa (maksimal 51): ")
	fmt.Scan(&_2311102177) // Ambil jumlah mahasiswa

	for i := 1; i <= *_2311102177; i++ { // Loop sebanyak mahasiswa yang ada
		fmt.Printf("Masukkan NIM mahasiswa ke-%d: ", i)
		fmt.Scan(&data[i].NIM) // Input NIM
		fmt.Printf("Masukkan nama mahasiswa ke-%d: ", i)
		fmt.Scan(&data[i].nama) // Input nama
		fmt.Printf("Masukkan nilai mahasiswa ke-%d: ", i)
		fmt.Scan(&data[i].nilai) // Input nilai
	}
}

// Fungsi buat cari nilai pertama anak mahasiswa dengan NIM tertentu
func cariNilaiPertama(data arrayMahasiswa, _2311102177 int, NIM string) (int, bool) {
	for i := 1; i <= _2311102177; i++ { // Loop semua mahasiswa
		if data[i].NIM == NIM { // Kalau ketemu NIM-nya
			return data[i].nilai, true // Balikin nilai pertama
		}
	}
	return 0, false // Kalau nggak ketemu, balikin 0 dan false
}

// Fungsi buat cari nilai paling besar mahasiswa dengan NIM tertentu
func cariNilaiTerbesar(data arrayMahasiswa, _2311102177 int, NIM string) (int, bool) {
	maxNilai := -1                      // Awalnya nilai terbesar kosong (pakai -1)
	found := false                      // Belum nemu siapa-siapa
	for i := 1; i <= _2311102177; i++ { // Loop semua mahasiswa
		if data[i].NIM == NIM { // Kalau NIM-nya cocok
			if data[i].nilai > maxNilai { // Kalau nilainya lebih besar
				maxNilai = data[i].nilai // Simpen nilai yang lebih besar
				found = true             // Udah nemu nilainya
			}
		}
	}
	return maxNilai, found // Balikin nilai terbesar (atau nggak ada)
}

func main() {
	var data arrayMahasiswa // Array buat nampung data mahasiswa
	var _2311102177 int     // Ini variabel jumlah mahasiswa
	var NIM string          // NIM mahasiswa yang mau dicari

	// Panggil fungsi buat masukin data mahasiswa
	inputMahasiswa(&data, &_2311102177)

	// Minta user masukin NIM yang mau dicari
	fmt.Print("Masukkan NIM yang ingin dicari: ")
	fmt.Scan(&NIM)

	// Cari nilai pertama
	nilaiPertama, foundPertama := cariNilaiPertama(data, _2311102177, NIM)
	if foundPertama {
		// Kalau ketemu, kasih tau nilai pertamanya
		fmt.Printf("Nilai pertama mahasiswa dengan NIM %s adalah: %d\n", NIM, nilaiPertama)
	} else {
		// Kalau nggak ketemu, bilang nggak ada
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan.\n", NIM)
	}

	// Cari nilai terbesar
	nilaiTerbesar, foundTerbesar := cariNilaiTerbesar(data, _2311102177, NIM)
	if foundTerbesar {
		// Kalau ketemu, kasih tau nilai terbesarnya
		fmt.Printf("Nilai terbesar mahasiswa dengan NIM %s adalah: %d\n", NIM, nilaiTerbesar)
	} else {
		// Kalau nggak ketemu, bilang nggak ada
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan.\n", NIM)
	}
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/2bb3c40c-9262-4aff-9abc-56f76c184e8c)

## Soal No 3

#### Source Code

```go
package main

import (
	"fmt"
	"strings"
)

// Konstanta untuk jumlah provinsi, ada 34 provinsi di Indonesia!
const nProv = 34

func main() {
	// Ini buat nyimpen nama-nama provinsi, populasi, dan angka pertumbuhannya
	var _2311102177 [nProv]string // Nama provinsi
	var PopProv [nProv]int        // Populasi provinsi
	var TumbuhProv [nProv]float64 // Angka pertumbuhan provinsi
	var ProvinsiDicari string     // Nama provinsi yang dicari

	// Yuk, masukkan data provinsi
	InputData(&_2311102177, &PopProv, &TumbuhProv)

	// Masukkan nama provinsi yang mau dicari
	fmt.Println("Masukkan nama provinsi yang mau dicari:")
	fmt.Scanln(&ProvinsiDicari)

	// Cari provinsi dengan pertumbuhan tercepat
	provTercepat := ProvinsiTercepat(_2311102177, TumbuhProv)
	fmt.Println("Provinsi dengan pertumbuhan tercepat adalah:", provTercepat)

	// Cari indeks provinsi yang dicari
	index := IndeksProvinsi(_2311102177, ProvinsiDicari)
	if index != -1 {
		// Kalau ketemu, kasih tau deh indeksnya
		fmt.Printf("Indeks provinsi %s adalah: %d\n", ProvinsiDicari, index)
	} else {
		// Kalau nggak ketemu, bilang aja nggak ketemu
		fmt.Printf("Provinsi %s nggak ditemukan.\n", ProvinsiDicari)
	}

	// Prediksi jumlah penduduk untuk provinsi dengan pertumbuhan di atas 2%
	fmt.Println("Prediksi jumlah penduduk untuk provinsi dengan pertumbuhan di atas 2%:")
	Prediksi(_2311102177, PopProv, TumbuhProv)
}

// Fungsi bayi buat input data nama provinsi, populasi, dan pertumbuhannya
func InputData(_2311102177 *[nProv]string, PopProv *[nProv]int, TumbuhProv *[nProv]float64) {
	fmt.Println("Masukkan data provinsi (Nama, Populasi, Pertumbuhan):")
	for i := 0; i < nProv; i++ {
		fmt.Printf("Data provinsi ke-%d:\n", i+1)
		// Masukkan data provinsi ke dalam array
		fmt.Scan(&_2311102177[i], &PopProv[i], &TumbuhProv[i])
	}
}

// Fungsi bayi buat cari provinsi dengan angka pertumbuhan tercepat
func ProvinsiTercepat(_2311102177 [nProv]string, TumbuhProv [nProv]float64) string {
	maxGrowth := TumbuhProv[0] // Anggap pertumbuhan pertama yang terbesar
	index := 0                 // Indeksnya provinsi pertama

	for i := 1; i < nProv; i++ {
		// Kalau ada yang lebih cepat, ganti deh
		if TumbuhProv[i] > maxGrowth {
			maxGrowth = TumbuhProv[i]
			index = i
		}
	}

	// Kembalikan nama provinsi dengan pertumbuhan tercepat
	return _2311102177[index]
}

// Fungsi bayi buat cari indeks provinsi berdasarkan nama
func IndeksProvinsi(_2311102177 [nProv]string, nama string) int {
	for i := 0; i < nProv; i++ {
		// Kalau ketemu, langsung balikin indeksnya
		if strings.EqualFold(_2311102177[i], nama) {
			return i
		}
	}
	// Kalau nggak ketemu, bilang -1 aja
	return -1
}

// Fungsi bayi buat prediksi populasi provinsi dengan pertumbuhan di atas 2%
func Prediksi(_2311102177 [nProv]string, PopProv [nProv]int, TumbuhProv [nProv]float64) {
	for i := 0; i < nProv; i++ {
		// Kalau pertumbuhannya lebih dari 2%, kita prediksi deh
		if TumbuhProv[i] > 2.0 {
			// Hitung prediksi populasi tahun depan
			prediksi := float64(PopProv[i]) * (1 + TumbuhProv[i]/100)
			// Tampilkan nama provinsi dan hasil prediksinya
			fmt.Printf("Provinsi %s: %.0f\n", _2311102177[i], prediksi)
		}
	}
}
```

#### Screenshoot Output

## Soal No 4

#### Source Code

```go
package main

import (
	"fmt"
	"sort"
)

const NMAX = 1000000 // Nih batas data yang bisa disimpan, maksimal 1 juta!

// Fungsi bayi buat hitung median
func median(_2311102177 []int) float64 {
	n := len(_2311102177) // Ukur panjangnya data yang ada
	if n%2 == 1 {
		// Kalau jumlah datanya ganjil, mediannya yang tengah aja
		return float64(_2311102177[n/2])
	}
	// Kalau jumlah datanya genap, mediannya rata-rata dua angka tengah
	return float64((_2311102177[n/2-1] + _2311102177[n/2])) / 2.0
}

func main() {
	var T []int // Ini wadahnya buat simpan angka-angka yang dibaca
	var x int   // Ini buat tampung angka yang bakal dimasukin

	// Nih kita mulai baca data dari input
	for {
		fmt.Scan(&x) // Baca angka dari input
		if x == -5313541 {
			// Kalau ketemu angka ini, berarti udah selesai, waktunya keluar
			break
		}
		if x == 0 {
			// Kalau dapet angka 0, kita cari median deh!
			if len(T) > 0 {
				sort.Ints(T)     // Urutkan dulu angka-angkanya biar rapi
				med := median(T) // Hitung mediannya
				// Kalau mediannya bulat, print tanpa koma
				if med == float64(int(med)) {
					fmt.Println(int(med)) // Print integer aja
				} else {
					fmt.Printf("%.1f\n", med) // Kalau enggak, print dengan 1 angka desimal
				}
			}
		} else {
			// Kalau bukan 0, simpen deh angkanya
			T = append(T, x)
		}
	}
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/6d5a70bb-55ea-4afb-84fb-7529014ee0f9)

## Soal No 5

#### Source Code

```go
package main

import (
	"fmt"
)

// Batas maksimal jumlah partai
const NMAX = 1000000

// Struct partai untuk nyimpen nama dan suara
type partai struct {
	nama  int // Nama partai
	suara int // Jumlah suara partai
}

// tipe tabPartai: array of partai, array besar NMAX
type tabPartai [NMAX]partai

func main() {
	// Deklarasi variabel
	var p tabPartai
	var _2311102177, x int // variabel _2311102177 digunakan untuk jumlah partai yang sudah dimasukin

	// Mulai hitung suara partai
	_2311102177 = 0 // Ini buat nyimpen jumlah partai yang udah dimasukin
	// Baca suara partai sampe ketemu -1, baru stop
	for {
		fmt.Scan(&x) // Baca suara partai
		if x == -1 { // Kalau -1, berhenti
			break
		}

		// Cek, partai udah ada atau belum
		pos := posisi(p, _2311102177, x)
		if pos == -1 { // Kalau belum ada, tambah partai baru
			p[_2311102177].nama = x
			p[_2311102177].suara = 1 // Suaranya mulai 1
			_2311102177++            // Jangan lupa, tambah jumlah partai yang dimasukin
		} else { // Kalau partainya udah ada, tambah suaranya
			p[pos].suara++
		}
	}

	// Urutkan suara dari yang banyak ke sedikit
	insertionSortDescending(p[:_2311102177], _2311102177)

	// Tampilkan suara partai
	for i := 0; i < _2311102177; i++ {
		// Tampilinnya sesuai format
		fmt.Printf("%d(%d) ", p[i].nama, p[i].suara)
	}
	fmt.Println()
}

// Fungsi ini buat cari, partai ada di mana
func posisi(t tabPartai, n int, nama int) int {
	// Cari partai yang punya nama ini
	for i := 0; i < n; i++ {
		if t[i].nama == nama {
			return i // Dapet, kembalikan posisi partainya
		}
	}
	return -1 // Kalau gak ketemu, return -1
}

// Fungsi ini buat urutin partai berdasarkan suara, urutannya dari banyak ke sedikit
func insertionSortDescending(arr []partai, n int) {
	// Start dari partai kedua
	for i := 1; i < n; i++ {
		j := i
		// Kalau suara partai ini lebih banyak dari yang sebelumnya, tuker tempatnya
		for j > 0 && arr[j-1].suara < arr[j].suara {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j-- // Terus cari yang lebih besar
		}
	}
}
```

#### Screenshoot Output

![image](https://github.com/user-attachments/assets/fc2e080f-7ca3-4526-8b4b-3e49a005a766)
