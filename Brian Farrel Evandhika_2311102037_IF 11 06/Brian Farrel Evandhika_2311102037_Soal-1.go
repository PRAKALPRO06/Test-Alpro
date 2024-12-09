// 2311102037 Brian Farrel Evandhika IF 11 06
package main

import "fmt"

type set [2022]int

func exist(T set, n_2311102037 int, val int) bool { //mengecek sebuah nilai pada array
	for i := 0; i < n_2311102037; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

func inputSet(T *set, n_2311102037 *int) { // memasukan nilai ke dalam himpunan hingga ditemukan duplikat pada himpunan tersebut
	var val int
	*n_2311102037 = 0
	fmt.Println("masukan himpunan bilangan bulat:")
	for {
		fmt.Scan(&val)
		if exist(*T, *n_2311102037, val) {
			break
		}
		T[*n_2311102037] = val
		(*n_2311102037)++
	}
}

func findIntersection(T1, T2 set, n_2311102037, m int, T3 *set, h *int) { //mencari irisan dua himpunan
	*h = 0
	for i := 0; i < n_2311102037; i++ {
		if exist(T2, m, T1[i]) {
			T3[*h] = T1[i]
			(*h)++
		}
	}
}

func printSet(T set, n_2311102037 int) { //menampilkan himpunan
	for i := 0; i < n_2311102037; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()
}

func main() { //fungsi utama untuk memanggil himpunan lainya
	var s1_2311102037, s2_2311102037, s3_2311102037 set
	var n1_2311102037, n2_2311102037, n3_2311102037 int
	inputSet(&s1_2311102037, &n1_2311102037)
	inputSet(&s2_2311102037, &n2_2311102037)
	findIntersection(s1_2311102037, s2_2311102037, n1_2311102037, n2_2311102037, &s3_2311102037, &n3_2311102037)
	fmt.Println("irisan himpunan:")
	printSet(s3_2311102037, n3_2311102037)
}
