package main

import "fmt"

type set_2311102226 [2022]int

func exist(T set_2311102226, n int, val int) bool {
	// Mengembalikan true apabila bilangan val ada di dalam array T
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

func inputSet(T *set_2311102226, n *int) {
	// Membaca input bilangan dan memasukkannya ke dalam array T
	var val int
	*n = 0
	for {
		fmt.Scan(&val)
		if exist(*T, *n, val) {
			break
		}
		T[*n] = val
		*n++
	}
}

func findIntersection(T1, T2 set_2311102226, n, m int, T3 *set_2311102226, k *int) {
	// Menghitung irisan antara T1 dan T2 dan menyimpannya di T3
	*k = 0
	for i := 0; i < n; i++ {
		if exist(T2, m, T1[i]) {
			if !exist(*T3, *k, T1[i]) {
				T3[*k] = T1[i]
				*k++
			}
		}
	}
}

func printSet(T set_2311102226, n int) {
	// Berfungsi menampilkan isi array T secara horizontal
	for i := 0; i < n; i++ {
		fmt.Print(T[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	var s1, s2, s3 set_2311102226
	var n1, n2, n3 int

	inputSet(&s1, &n1)
	inputSet(&s2, &n2)
	findIntersection(s1, s2, n1, n2, &s3, &n3)
	printSet(s3, n3)
}
