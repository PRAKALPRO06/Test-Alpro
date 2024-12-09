package main

import "fmt"

type set [2022]int

// Fungsi untuk mengecek apakah elemen sudah ada dalam array
func exist_2311102018(T set, n int, val int) bool {
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

// Fungsi untuk mengisi array dengan bilangan unik
func inputSet(T *set, n *int) {
	var val int
	*n = 0
	for {
		fmt.Scan(&val)
		if exist_2311102018(*T, *n, val) {
			break
		}
		T[*n] = val
		(*n)++
	}
}

// Fungsi untuk mencari irisan dari dua array
func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	*h = 0
	for i := 0; i < n; i++ {
		if exist_2311102018(T2, m, T1[i]) && !exist_2311102018(*T3, *h, T1[i]) {
			T3[*h] = T1[i]
			(*h)++
		}
	}
}

// Fungsi untuk mencetak array secara horizontal
func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(T[i])
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