package main

import (
	"fmt"
)

// Nama: Arjun Ahmad Santoso
// NIM: 231102211

const n_max int = 100

// Fungsi untuk menemukan bilangan dalam suatu array
func find(x_231102211 int, arr [n_max]int, n int) bool {
	for i := 0; i < n; i++ {
		if x_231102211 == arr[i] {
			return true
		}
	}
	return false
}

func find_irisan(arr1 [n_max]int, n1 int, arr2 [n_max]int, n2 int, arrIrisan [n_max]int, nIrisan int) {
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if arr1[i] == arr2[j] {
				arrIrisan[nIrisan] = arr1[i]
				nIrisan++
			}
		}
	}
}

func main() {

	arr1 := [n_max]int{}
	n1 := 0
	arr2 := [n_max]int{}
	n2 := 0
	arrIrisan := [n_max]int{}
	nIrisan := 0

	for i := 0; i < n_max; i++ {
		fmt.Println("Masukkan bilangan pada himpunan pertama: ")

		var x int
		fmt.Scan(&x)

		if find(x, arr1, n1) {
			break
		} else {
			arr1[i] = x
			n1++
		}
	}

	fmt.Print("\n")

	for i := 0; i < n_max; i++ {
		fmt.Println("Masukkan bilangan pada himpunan kedua: ")

		var x int
		fmt.Scan(&x)

		if find(x, arr2, n2) {
			break
		} else {
			arr2[i] = x
			if find(x, arr1, n1) {
				arrIrisan[nIrisan] = x
				nIrisan++
			}
			n2++
		}
	}

	fmt.Println("Irisan: ")
	for i := 0; i < nIrisan; i++ {
		fmt.Print(arrIrisan[i])
	}
}
