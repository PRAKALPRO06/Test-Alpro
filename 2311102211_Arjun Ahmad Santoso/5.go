package main

import "fmt"

type Partai struct {
	nama  int
	suara int
}

const nMax int = 1000000

type arrPartai [nMax]Partai

func posisi(arr arrPartai, nPartai int, nama_2311102211 int) int {
	for i := 0; i < nPartai; i++ {
		if nama_2311102211 == arr[i].nama {
			return i
		}
	}
	return -1
}

func InputData(arr arrPartai, nPartai *int) {
	fmt.Println("Masukkan nama partai")
	var nama int
	for i := 0; i < nMax; i++ {
		fmt.Scan(&nama)
		if nama == -1 {
			return
		}
		if pos := posisi(arr, nMax, nama); pos != -1 {
			arr[pos].suara++
		} else {
			arr[i].nama = nama
			*nPartai++
		}
	}
}

func InsertionSort(arr arrPartai, nPartai *int) {
	for i := 0; i < *nPartai; i++ {
		fmt.Println(arr[i].nama, "(", arr[i].suara)
	}
}

func main() {
	var arr arrPartai
	nPartai := 0

	InputData(arr, &nPartai)

	fmt.Println("\n Data Partai: ")
	for i := 0; i < nPartai; i++ {
		fmt.Println(arr[i].nama, "(", arr[i].suara)
	}

}
