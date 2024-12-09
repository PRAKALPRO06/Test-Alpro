package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Fungsi untuk pencarian median
func cariMedian_2311102327(angka []int) float64 {
	n := len(angka)
	sort.Ints(angka)
	if n%2 == 0 {
		return float64(angka[n/2-1]+angka[n/2]) / 2.0
	}
	return float64(angka[n/2])
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input deretan angka (akhiri dengan -5313541 kalau mau end):")
	scanner.Scan()
	masukan := scanner.Text()
	
	angkaStr := strings.Fields(masukan)
	var daftarAngka []int

	for _, val := range angkaStr {
		angka, _ := strconv.Atoi(val)
		if angka == -5313541 {
			break
		}
		if angka != 0 {
			daftarAngka = append(daftarAngka, angka)
		} else {
			median := cariMedian_2311102327(daftarAngka)
			fmt.Printf("Median: %.1f\n", median)
		}
	}
}