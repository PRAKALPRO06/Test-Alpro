package main

import (
	"fmt"
)

const nMax = 51

type Mahasiswa struct {
    NIM   string
    Nama  string
    Nilai int
}

type ArrayMahasiswa [nMax]Mahasiswa

//function untuk mencari nilai menggunakan perbandingan
func cariNilaiPertamaDanTerbesar(arr ArrayMahasiswa, jumlahData int, nim string) (int, int) {
    nilaiPertama := -1
    nilaiTerbesar := -1

    for i := 0; i < jumlahData; i++ {
        if arr[i].NIM == nim {
            if nilaiPertama == -1 {
                nilaiPertama = arr[i].Nilai
                nilaiTerbesar = arr[i].Nilai
            } else {
                if arr[i].Nilai > nilaiTerbesar {
                    nilaiTerbesar = arr[i].Nilai
                }
            }
        }
    }

    return nilaiPertama, nilaiTerbesar
}

func main() {
    var dataMahasiswa ArrayMahasiswa
    var jumlahData int

    fmt.Print("Masukkan jumlah data mahasiswa: ")
    fmt.Scan(&jumlahData)

    if jumlahData > nMax {
        fmt.Println("Jumlah data melebihi kapasitas maksimal.")
        return
    }

    // Input data mahasiswa
    for i := 0; i < jumlahData; i++ {
        fmt.Printf("Masukkan NIM mahasiswa ke-%d: ", i+1)
        fmt.Scan(&dataMahasiswa[i].NIM)
        fmt.Printf("Masukkan nama mahasiswa ke-%d: ", i+1)
        fmt.Scan(&dataMahasiswa[i].Nama)
        fmt.Printf("Masukkan nilai mahasiswa ke-%d: ", i+1)
        fmt.Scan(&dataMahasiswa[i].Nilai)
    }

    var nim_2311102226 string
    fmt.Print("Masukkan NIM yang dicari: ")
    fmt.Scan(&nim_2311102226)

    nilaiPertama, nilaiTerbesar := cariNilaiPertamaDanTerbesar(dataMahasiswa, jumlahData, nim_2311102226)

    if nilaiPertama == -1 {
        fmt.Println("Mahasiswa dengan NIM tersebut tidak ditemukan.")
    } else {
        fmt.Printf("Nilai pertama mahasiswa dengan NIM %s adalah %d\n", nim_2311102226, nilaiPertama)
        fmt.Printf("Nilai terbesar mahasiswa dengan NIM %s adalah %d\n", nim_2311102226, nilaiTerbesar)
    }
}

