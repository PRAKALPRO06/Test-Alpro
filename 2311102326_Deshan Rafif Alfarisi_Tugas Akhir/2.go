package main

import"fmt"

//Struck nama, nim, nilai

type mahasiswa struct{
	nama, nim2311102326 string
	nilai float64
}

//tipe data array mahasiswa max 2023
type arrMhs[2023] mahasiswa



func main(){
	var n2311102326 int 
	var dataMhs arrMhs

	//meminta input jumlah mahasiswa
	fmt.Print("Masukkan jumlah mahasiswa : ")
	fmt.Scan(&n2311102326)

	//validasi jumlah mahasiswa yang dimasukkan 
	if n2311102326 < 1 || n2311102326 >2023{
		fmt.Println("Jumlah mahasiswa harus antara 1 dan 2023")
		return
	}
	//Mengisi data mahasiswa
	for i :=0; i <n2311102326; i++{
		fmt.Printf("\nMasukkan data mahasiswa ke- %d\n", i+1)
		fmt.Print("Nama : ")
		fmt.Scan(&dataMhs[i].nama)
		fmt.Print("NIM : ")
		fmt.Scan(&dataMhs[i].nim2311102326)
		fmt.Print("Nilai: ")
		fmt.Scan(&dataMhs[i].nilai)
	}

	//nilai terbesar
	nilaiTerbesar := NilaiTertinggi(dataMhs, n2311102326)
}
