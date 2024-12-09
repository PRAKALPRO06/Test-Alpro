package main

import "fmt"

// di fungsi ini untuk menginput jumlah element
func input_2311102241() ([]int, int) {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	set := make([]int, 0)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		if !exist(set, x) {
			set = append(set, x)
		}
	}
	return set, len(set)
}

func exist(set []int, val int) bool {
	for _, v := range set {
		if v == val {
			return true
		}
	}
	return false
}

func find_2311102241(set1, set2 []int) []int {
	intersection := make([]int, 0)
	for _, v := range set1 {
		if exist(set2, v) {
			intersection = append(intersection, v)
		}
	}
	return intersection
}

// di fungsi ini untuk menampilkan hasil print set
func printSet(set []int) {
	for i, v := range set {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func main() {
	set1, _ := input_2311102241()
	set2, _ := input_2311102241()

	intersection := find_2311102241(set1, set2)

	println("Output set")
	printSet(intersection)
}
