package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan baris pertama_2311102249:")
	scanner.Scan()
	line1 := scanner.Text()

	fmt.Println("Masukkan baris kedua_2311102249:")
	scanner.Scan()
	line2 := scanner.Text()

	// Konversi masukan string ke slice of int
	input1 := parseInput(line1)
	input2 := parseInput(line2)

	// Menghapus duplikasi pada setiap baris
	set1 := removeDuplicates(input1)
	set2 := removeDuplicates(input2)

	// Mencari irisan (intersection)
	intersection := findIntersection(set1, set2)

	// Menampilkan keluaran
	fmt.Println("Keluaran:", intersection)
}

// Fungsi untuk menghapus duplikasi dari slice
func removeDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, num := range arr {
		if _, exists := seen[num]; exists {
			break
		}
		seen[num] = true
		result = append(result, num)
	}

	return result
}

// Fungsi untuk mencari irisan dua himpunan
func findIntersection(set1, set2 []int) []int {
	intersection := []int{}
	seen := make(map[int]bool)

	// Tandai elemen yang ada di set2
	for _, num := range set2 {
		seen[num] = true
	}

	// Cari elemen yang ada di kedua himpunan
	for _, num := range set1 {
		if seen[num] {
			intersection = append(intersection, num)
		}
	}

	return intersection
}

// Fungsi untuk mengonversi input string ke slice of int
func parseInput(input string) []int {
	parts := strings.Fields(input)
	result := []int{}

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			result = append(result, num)
		}
	}

	return result
}
