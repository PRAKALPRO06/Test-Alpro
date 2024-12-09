package main

import "fmt"

type set [2022]int

// fungsi untuk mengembalikan nilai true apabilang bilangan yang ada di array T berisi sejumlah n
func exist_2311102010(T set, n, val int) bool {
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

// fungsi untuk menambahkan nilai  ke array T dan memberhentikan inputan dengan diakhiri -1
func inputSet(T *set, n *int) {
	var val int
	*n = 0
	for {
		fmt.Scan(&val)
		if val == -1 {
			break
		}
		if !exist_2311102010(*T, *n, val) {
			T[*n] = val
			*n++
		}
	}
}

// fungsi untuk mencari irisan antar baris yang ada pada array T
func findIntersection(T1, T2 set, n1, n2 int, T3 *set, n3 *int) {
	*n3 = 0
	for i := 0; i < n1; i++ {
		if exist_2311102010(T2, n2, T1[i]) {
			T3[*n3] = T1[i]
			*n3++
		}
	}
}

// fungsi untuk mencetak nilai yang ada pada irisan
func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", T[i])
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