package main

import "fmt"

type set [2022]int

// Fungsi untuk memeriksa apakah nilai sudah ada dalam array
func exist_2311102327(T set, n int, val int) bool {
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

// Fungsi untuk menerima input himpunan
func inputSet(T *set, n *int) {
	var val int
	*n = 0
	for {
		fmt.Scan(&val)
		if exist_2311102327(*T, *n, val) {
			break
		}
		(*T)[*n] = val
		*n++
	}
}

// Fungsi untuk menemukan irisan dua himpunan
func findIntersection(T1, T2 set, n1, n2 int, T3 *set, n3 *int) {
	*n3 = 0
	for i := 0; i < n1; i++ {
		if exist_2311102327(T2, n2, T1[i]) {
			if !exist_2311102327(*T3, *n3, T1[i]) {
				(*T3)[*n3] = T1[i]
				(*n3)++
			}
		}
	}
}

// Fungsi untuk mencetak isi array
func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
	fmt.Println()
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int

	inputSet(&s1, &n1)
	inputSet(&s2, &n2)
	findIntersection(s1, s2, n1, n2, &s3, &n3)
	printSet(s3, n3)
}