package main

import (
	"fmt"
)

type Provinsi struct {
	nama              string
	jumlah_populasi   int
	angka_pertumbuhan int
}

const nMax int = 34

type arrProvinsi [nMax]Provinsi

func InputData(arr arrProvinsi) {
	for i := 0; i < nMax; i++ {
		fmt.Scan(&arr[i].nama, &arr[i].jumlah_populasi, &arr[i].angka_pertumbuhan)
	}
}
func ProvinsiTercepat(arr arrProvinsi) Provinsi {
	provinsi_tercepat := arr[0]
	for i := 0; i < nMax; i++ {
		if arr[i].angka_pertumbuhan > provinsi_tercepat.angka_pertumbuhan {
			provinsi_tercepat = arr[i]
		}
	}
	return provinsi_tercepat
}
func TampilkanPrediksi(arr arrProvinsi) {
	for i := 0; i < nMax; i++ {
		fmt.Println(arr[i].nama, ", prediksi jumlah populasinya di tahun depan: ", arr[i].jumlah_populasi*arr[i].angka_pertumbuhan)
	}
}
func IndexProvinsi(arr arrProvinsi, nama string) int {
	for i := 0; i < nMax; i++ {
		if arr[i].nama == nama {
			return i
		}
	}
	return -1
}

func main() {

	var arr arrProvinsi

	fmt.Println("Masukkan nama, jumlah jumlah_populasi dan angka pertumbuhan Provinsi di setiap baris (dipisahkan dengan spasi): ")
	InputData(arr)

	fmt.Print("Masukkan nama Provinsi yang ingin dicari: ")
	var nama_cari string
	fmt.Scan(&nama_cari)

	fmt.Println("Provinsi dengan angka pertumbuhan tercepat adalah: ", ProvinsiTercepat(arr))
	fmt.Println("Index provinsi yang dicari adalah: ", IndexProvinsi(arr, nama_cari))
	fmt.Println("Prediksi Jumlah Populasi setiap provinsi adalah: ")
	TampilkanPrediksi(arr)

}
