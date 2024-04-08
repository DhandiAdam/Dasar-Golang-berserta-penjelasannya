package main

import "fmt"

func main() {
	//Nilai Const adalah nilai tetap dan tidak bisa diubah variablenya
	//Jadi harus Consisten saat memakai tipe data tersebut
	//Const gabungan

	const (
		nama         string = "Dhandi adam"
		namaBelakang        = "Adam"
	)

	fmt.Println("Nama Lengkap = ", nama)
	fmt.Println("Nama Belakang = ", namaBelakang)

}
