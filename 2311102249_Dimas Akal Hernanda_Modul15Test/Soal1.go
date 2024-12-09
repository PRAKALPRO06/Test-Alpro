package main

import "fmt"

type set []int

//mengecek nilai
func exist(T set, n int, val int) bool {
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

// inputSet membaca nilai array
func inputSet(T *set, n *int) {
	var num int
	*n = 0

	for {
		fmt.Scan(&num)
		// mengecek nilai
		if exist(*T, *n, num) {
			break
		}
		(*T)[*n] = num
		*n++
	}
}

func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	*h = 0
	for i := 0; i < n; i++ {
		if exist(T2, m, T1[i]) {
			if !exist(*T3, *h, T1[i]) {
				(*T3)[*h] = T1[i]
				*h++
			}
		}
	}
}

// Memasukan nilai elemen
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
	var s1, s2, s3 set = make(set, 100), make(set, 100), make(set, 100)
	var n1, n2, n3 int

	inputSet(&s1, &n1)

	inputSet(&s2, &n2)

	findIntersection(s1, s2, n1, n2, &s3, &n3)

	printSet(s3, n3)
}
