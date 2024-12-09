package main

import (
	"fmt"
	"strings"
)

const maxProv = 34

func inputData() (int, [maxProv]string, [maxProv]int, [maxProv]float64) {
	var n int
	var namaProv [maxProv]string
	var popProv [maxProv]int
	var tumbuhProv [maxProv]float64

	fmt.Println("Masukkan jumlah provinsi (maksimal 34):")
	fmt.Scan(&n)
	if n > maxProv {
		fmt.Printf("Jumlah provinsi melebihi batas maksimum (%d). Menggunakan nilai maksimum.\n", maxProv)
		n = maxProv
	} else if n <= 0 {
		fmt.Println("Jumlah provinsi harus lebih dari 0.")
		return 0, namaProv, popProv, tumbuhProv
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan nama provinsi ke-%d: ", i+1)
		fmt.Scan(&namaProv[i])
		namaProv[i] = strings.TrimSpace(namaProv[i]) // Menghapus spasi berlebih

		fmt.Printf("Masukkan populasi provinsi ke-%d (juta): ", i+1)
		fmt.Scan(&popProv[i])
		if popProv[i] <= 0 {
			fmt.Println("Populasi harus lebih dari 0. Masukkan ulang.")
			i-- // Ulangi input untuk provinsi ini
			continue
		}

		fmt.Printf("Masukkan pertumbuhan penduduk provinsi ke-%d (%%): ", i+1)
		fmt.Scan(&tumbuhProv[i])
		if tumbuhProv[i] < 0 {
			fmt.Println("Pertumbuhan tidak boleh negatif. Masukkan ulang.")
			i-- // Ulangi input untuk provinsi ini
			continue
		}
	}

	return n, namaProv, popProv, tumbuhProv
}

func provinsiTerbesar(popProv [maxProv]int, n int) int {
	maxIndex := 0
	for i := 1; i < n; i++ {
		if popProv[i] > popProv[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

func prediksiPenduduk(popProv [maxProv]int, tumbuhProv [maxProv]float64, n int) [maxProv]int {
	var prediksi [maxProv]int
	for i := 0; i < n; i++ {
		prediksi[i] = popProv[i] + int(float64(popProv[i])*tumbuhProv[i]/100)
	}
	return prediksi
}

func provinsiPertumbuhanTinggi(tumbuhProv [maxProv]float64, n int) []int {
	var result []int
	for i := 0; i < n; i++ {
		if tumbuhProv[i] > 2.0 {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	n, namaProv, popProv, tumbuhProv := inputData()
	if n == 0 {
		return
	}

	// Provinsi dengan populasi terbesar
	maxIndex := provinsiTerbesar(popProv, n)
	fmt.Printf("\nProvinsi dengan populasi terbesar: %s\n", namaProv[maxIndex])

	// Prediksi populasi setelah pertumbuhan
	prediksi := prediksiPenduduk(popProv, tumbuhProv, n)
	fmt.Println("\nPrediksi populasi setelah pertumbuhan:")
	for i := 0; i < n; i++ {
		fmt.Printf("%s: %d juta\n", namaProv[i], prediksi[i])
	}

	// Provinsi dengan pertumbuhan di atas 2%
	tinggiPertumbuhan := provinsiPertumbuhanTinggi(tumbuhProv, n)
	fmt.Println("\nProvinsi dengan pertumbuhan di atas 2%:")
	if len(tinggiPertumbuhan) == 0 {
		fmt.Println("Tidak ada provinsi dengan pertumbuhan di atas 2%.")
	} else {
		for _, index := range tinggiPertumbuhan {
			fmt.Printf("%s\n", namaProv[index])
		}
	}
}3