package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	// Memeriksa apakah nilai val ada di dalam array T yang berisi n bilangan bulat
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

func inputSet(T *set, n *int) {
	// I.S. data himpunan telah siap pada piranti masukan
	// F.S. array T berisi sejumlah n bilangan bulat yang berasal dari masukan
	// (masukan berakhir apabila bilangan ada yang duplikat, atau array penuh)
	var input int
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}
		if !exist(*T, *n, input) {
			T[*n] = input
			*n++
		} else {
			break // Berhenti jika ada duplikat
		}
	}
}

func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	// I.S. terdefinisi himpunan T1 dan T2 yang berisi sejumlah n dan m anggota himpunan
	// F.S. himpunan T3 berisi sejumlah h bilangan bulat yang merupakan irisan dari himpunan T1 dan T2
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if T1[i] == T2[j] && !exist(*T3, *h, T1[i]) {
				T3[*h] = T1[i]
				*h++
				break
			}
		}
	}
}

func printSet(T set, n int) {
	// I.S. terdefinisi sebuah himpunan T yang berisi sejumlah n bilangan bulat
	// F.S. menampilkan isi array T secara horizontal (dipisahkan oleh spasi)
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
