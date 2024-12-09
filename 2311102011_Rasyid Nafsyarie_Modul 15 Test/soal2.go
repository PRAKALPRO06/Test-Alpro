package main
import "fmt"
const nMax = 51
type mahasiswa struct {
	NIM string
	nama string
	nilai int
}

type arrayMahasiswa [nMax] mahasiswa


//fungsi ini yaitu proses input mahasiswa beserta data diri lainnya
func inputMahasiswa (T *arrayMahasiswa, N *int) {
	fmt.Print("Masukkan Jumlah data Mahasiswa : ")
	fmt.Scan(N)
	for i := 0; i < *N; i++ {
		fmt.Printf("data mahasiswa ke-%d\n", i+1)
		fmt.Print("Masukkan NIM : ")
		fmt.Scan(&T[i].NIM)
		fmt.Print("Masukkan Nama : ")
		fmt.Scan(&T[i].nama)
		fmt.Print("Masukkan Nilai : ")
		fmt.Scan(&T[i].nilai)
	}
}

//fungsi ini untuk mencari sebuah nilai pertama dari mahasiswa yang sudah diinputkan tadi
func cariNilai (T arrayMahasiswa, N int, nim_2311102011 string) int {
	for i := 0; i < N; i++ {
		if T[i].NIM == nim_2311102011 {
			return T[i].nilai
		}
	}
	return -1
}

//kemudian fungsi nilai terbesar ini untuk menemukan nilai dari inputan tadi yang paling besar menggunakan for untuk melakukan pengecekan dan percabangan
func cariNilaiTerbesar(T arrayMahasiswa, N int, nim_2311102011 string) int{
	maxNilai := -1
	found := false
	for i := 0; i < N; i++ {
		if T[i].NIM == nim_2311102011 {
			found = true
			if T[i].nilai > maxNilai{
				maxNilai = T[i].nilai 	
			}
		}
	}
	if found {
		return maxNilai
	}
	return -1
}

func main() {
	var dataMahasiswa arrayMahasiswa
	var jumlahData int
	var nim_2311102011 string
	inputMahasiswa(&dataMahasiswa, &jumlahData)

	fmt.Print("Masukkan NIM untuk mencari nilai pertama : ")
	fmt.Scan(&nim_2311102011)
	nilaiPertama := cariNilai(dataMahasiswa, jumlahData, nim_2311102011)
	if nilaiPertama != -1 {
		fmt.Printf("Nilai pertama untuk NIM %s adalah %d\n", nim_2311102011, nilaiPertama)
	} else {
		fmt.Printf("Data dengan NIM %s tidak ditemukan\n", nim_2311102011)
	}
	fmt.Print("Masukkan NIM untuk mencari nilai terbesar: ")
	fmt.Scan(&nim_2311102011)
	nilaiTerbesar := cariNilaiTerbesar(dataMahasiswa, jumlahData, nim_2311102011)
	if nilaiTerbesar != -1 {
		fmt.Printf("Nilai terbesar untuk NIM %s adalah %d\n", nim_2311102011, nilaiTerbesar)
	} else {
		fmt.Printf("Data dengan NIM %s tidak ditemukan\n", nim_2311102011)
	}
}