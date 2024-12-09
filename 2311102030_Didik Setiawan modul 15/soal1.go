package main

import (
	"fmt"
	"strconv"
	"strings"
)

// fungsi untuk menginputkan dan memeriksa inputan
func inputSet() map[int]struct{} {
	var input string
	fmt.Scanln(&input)
	set := make(map[int]struct{})
	for _, val := range strings.Fields(input) {
		if num, err := strconv.Atoi(val); err == nil {
			set[num] = struct{}{}
		}
	}
	return set
}

// mengcari inputan
func findIntersection(set1, set2 map[int]struct{}) []int {
	var intersection []int
	for num := range set1 {
		if _, exists := set2[num]; exists {
			intersection = append(intersection, num)
		}
	}
	return intersection
}

func main() {
	//pengguna meninputkan
	fmt.Println("Masukkan himpunan pertama:")
	set1 := inputSet()

	fmt.Println("Masukkan himpunan kedua:")
	set2 := inputSet()

	//mengcetak
	fmt.Println("Irisan himpunan:", findIntersection(set1, set2))
}
