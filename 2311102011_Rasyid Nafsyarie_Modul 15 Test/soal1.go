package main
import "fmt"

type set[2022]int

//fungsi ini untuk melakukan pengecekan apakah nilainya sudah dalam array 
func exist (T set, n int, val int) bool {
	for i := 0; i < n; i++ {
		if T[i] == val {
			return true
		}
	}
	return false
}

//fungsi ini untuk mengisi nilai array
func inputSet (T *set, n *int) {
	var val int
	*n = 0
	for {
		fmt.Scan(&val)
		if exist(*T, *n, val) {
			break
		}
		T[*n] = val
		(*n)++
	}
}

//fungsi ini untuk mencari nilai array yang terduplikat
func findIntersection (T1, T2 set, n, m int, T3 *set, h *int) {
	*h = 0
	for i := 0; i < n; i++ {
		if exist(T2, m, T1[i]) && !exist(*T3, *h, T1[i]) {
			T3[*h] = T1[i]
			(*h)++
		}
	}
}

//fungsi ini untuk mencetak nilai array secara horizontal
func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(T[i])
	}
	fmt.Println()
}

func main () {
	var s1_2311102011, s2, s3 set
	var n1, n2, n3 int
	inputSet(&s1_2311102011, &n1)
	inputSet(&s2, &n2)
	findIntersection(s1_2311102011, s2,n1,n2, &s3, &n3,)
	printSet(s3,n3)
}