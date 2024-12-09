package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Membaca input dari pengguna
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan elemen baris pertama (pisahkan dengan spasi):")
	scanner.Scan()
	line1 := scanner.Text()

	fmt.Println("Masukkan elemen baris kedua (pisahkan dengan spasi):")
	scanner.Scan()
	line2 := scanner.Text()

	// Memproses input ke dalam himpunan unik
	set1 := toSet(line1)
	set2 := toSet(line2)

	// Mencari irisan kedua himpunan
	intersection := hanif2311102266(set1, set2)

	// Menampilkan hasil
	fmt.Println("Irisan dari himpunan:", intersection)
}

// Fungsi untuk mengubah input string menjadi himpunan unik
func toSet(input string) map[int]struct{} {
	elements := strings.Split(input, " ")
	set := make(map[int]struct{})

	for _, elem := range elements {
		if num, err := strconv.Atoi(elem); err == nil {
			set[num] = struct{}{}
		}
	}

	return set
}

// Fungsi untuk mencari irisan dua himpunan
func hanif2311102266(set1, set2 map[int]struct{}) []int {
	intersection := []int{}

	for key := range set1 {
		if _, exists := set2[key]; exists {
			intersection = append(intersection, key)
		}
	}

	return intersection
}
