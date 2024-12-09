package main

import "fmt"

const nProv = 34
type NamaProv [nProv] string
type PopProv [nProv] int
type TumbuhProv [nProv] float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv){
	for i := 0; i < nProv; i++ {
		fmt.Scan(prov, pop, tumbuh)
	}
	fmt.Println("Nama Provinsi yang dicari:")
	fmt.Scan(prov)
}

func ProvinsiTercepat(tumbuh TumbuhProv){

}

func Prediksi(){

}

func IndeksProvinsi(){

}


func main(){
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv

	InputData(&prov, &pop, &tumbuh)


}