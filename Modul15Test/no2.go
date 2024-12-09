package main

import "fmt"

const MAX_SISWA = 51
const MAX_NILAI = 10

type Student struct {
	NIM   string
	Nama  string
	Nilai []int
}

func main() {
	var students [MAX_SISWA]Student
	var hitungSiswa int

	// memasukkan jumlah siswa
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&hitungSiswa)

	// memasukkan detail siswa
	for i := 0; i < hitungSiswa; i++ {
		fmt.Printf("\nMahasiswa ke-%d\n", i+1)
		fmt.Print("NIM: ")
		fmt.Scan(&students[i].NIM)
		fmt.Print("Nama: ")
		fmt.Scan(&students[i].Nama)

		// memasukkan nilai lebih untuk setiap siswa
		var jumlahNilai int
		fmt.Print("Masukkan jumlah nilai: ")
		fmt.Scan(&jumlahNilai)

		students[i].Nilai = make([]int, jumlahNilai)
		for j := 0; j < jumlahNilai; j++ {
			fmt.Printf("Masukkan Nilai ke-%d: ", j+1)
			fmt.Scan(&students[i].Nilai[j])
		}
	}

	// mencari siswa
	var cariNIM_2311102273 string
	fmt.Print("\nMasukkan NIM yang dicari: ")
	fmt.Scan(&cariNIM_2311102273)

	// menemukan nilai pertama dan tertinggi
	var menemukanSiswa []Student

	// Cari semua siswa dengan NIM yang sama
	for i := 0; i < hitungSiswa; i++ {
		if students[i].NIM == cariNIM_2311102273 {
			menemukanSiswa = append(menemukanSiswa, students[i])
		}
	}

	// Tampilkan hasil pencarian
	if len(menemukanSiswa) > 0 {
		fmt.Printf("Ditemukan %d mahasiswa dengan NIM %s:\n", len(menemukanSiswa), cariNIM_2311102273)

		for _, student := range menemukanSiswa {
			fmt.Printf("\nNama: %s\n", student.Nama)

			// Tampilkan semua nilai
			fmt.Println("Daftar Nilai:")
			for j, grade := range student.Nilai {
				fmt.Printf("Nilai ke-%d: %d\n", j+1, grade)
			}

			// Cari nilai pertama dan tertinggi
			if len(student.Nilai) > 0 {
				nilaiPertama := student.Nilai[0]
				nilaiTertinggi := findHighestGrade(student.Nilai)

				fmt.Printf("Nilai pertama: %d\n", nilaiPertama)
				fmt.Printf("Nilai tertinggi: %d\n", nilaiTertinggi)
			}
		}
	} else {
		fmt.Printf("Mahasiswa dengan NIM %s tidak ditemukan\n", cariNIM_2311102273)
	}
}

// Fungsi untuk mencari nilai tertinggi
func findHighestGrade(nilai []int) int {
	if len(nilai) == 0 {
		return -1
	}

	tertinggi := nilai[0]
	for _, grade := range nilai {
		if grade > tertinggi {
			tertinggi = grade
		}
	}
	return tertinggi
}
