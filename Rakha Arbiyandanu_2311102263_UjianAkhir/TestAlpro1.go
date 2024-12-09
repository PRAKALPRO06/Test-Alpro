
package main

import "fmt"

func bacaHimpunan_2311102263() map[int]struct{} {//Untuk menyimpan data pada setiap bilangan yang diinputkan dan membaca data bilangan duplikat atau tidaknya
	himpunan := make(map[int]struct{})
	for {
		var bilangan int
		fmt.Print("Masukkan bilangan : ")
		_, err := fmt.Scan(&bilangan)
		if err != nil {
			fmt.Println("Masukkan bilangan bulat.")
			continue
		}

		if _, double := himpunan[bilangan]; double {
			fmt.Printf("%d duplikat, masukan berhenti.\n", bilangan)
			break
		}
		himpunan[bilangan] = struct{}{}
	}

	return himpunan
}

func main() {//fungsi ini adalah fungsi pertama untuk menjalankan program seperti membaca data,cek data dll
	fmt.Println("Masukkan bilangan untuk himpunan pertama:")
	himpunan1 := bacaHimpunan_2311102263()

	fmt.Println("\nMasukkan bilangan untuk himpunan kedua:")
	himpunan2 := bacaHimpunan_2311102263()


	irisan := make(map[int]struct{})
	for bilangan := range himpunan1 {
		if _, exists := himpunan2[bilangan]; exists {
			irisan[bilangan] = struct{}{}
		}
	}

	fmt.Println("\nHimpunan pertama:", keys(himpunan1))
	fmt.Println("Himpunan kedua:", keys(himpunan2))
	fmt.Println("Irisan kedua himpunan:", keys(irisan))
}

func keys(m map[int]struct{}) []int {//fungsi ini digunakan untuk membuat struct digubakan untuk membagi bilangan dalam setiap baris
	result := []int{}
	for k := range m {
		result = append(result, k)
	}
	return result
}
