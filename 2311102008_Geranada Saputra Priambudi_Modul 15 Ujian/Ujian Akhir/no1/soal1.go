// Geranada Saputra Priambudi
// 2311102008
// IF-11-06
package main

import "fmt"

// Definisikan tipe data set sebagai array dengan ukuran max 2022
type set [2022]int

// Fungsi untuk memeriksa apakah suatu nilai ada dalam set
func exist(T set, n_2311102008 int, val int) bool {
	var i int = 0
	var status bool = false
	for i < n_2311102008 && !status {
		status = T[i] == val
		i++
	}
	return status
}

// Fungsi untuk memasukkan data ke dalam set
func inputSet(T *set, n_2311102008 *int) {
	*n_2311102008 = 0
	var bilangan int
	fmt.Scan(&bilangan)
	for *n_2311102008 < 2022 && !exist(*T, *n_2311102008, bilangan) {
		T[*n_2311102008] = bilangan
		(*n_2311102008)++
		fmt.Scan(&bilangan)
	}
}

// Fungsi untuk mencari irisan dari dua himpunan
func findIntersection(T1, T2 set, n_2311102008, m int, T3 *set, h *int) {
	var j int = 0
	*h = 0
	for j < n_2311102008 {
		if exist(T2, m, T1[j]) {
			T3[*h] = T1[j]
			(*h)++
		}
		j++
	}
}

// Fungsi untuk mencetak set
func printSet(T set, n_2311102008 int) {
	for i := 0; i < n_2311102008; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int

	// Input set pertama
	inputSet(&s1, &n1)
	// Input set kedua
	inputSet(&s2, &n2)

	// Mencari irisan kedua set
	findIntersection(s1, s2, n1, n2, &s3, &n3)

	// Menampilkan hasil irisan
	printSet(s3, n3)
}
