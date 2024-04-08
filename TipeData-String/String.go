package main

import "fmt"

func main() {
	var nama string
	nama = "Dhandi"
	//Mendeklarasi variabel dengan menggunakan := tanpa harus membuat var nama string
	Nama_Belakang := "Adam"
	//Mendeklarasi variabel function tanpa harus mendefinisikan func nama_variabel
	Nama_Lengkap := func(NamaDepan, NamaBelakang string) string {
		return NamaDepan + " " + NamaBelakang
	}
	fmt.Println("Nama Depan = ", nama)
	fmt.Println("Nama Belakang = ", Nama_Belakang)
	fmt.Println("Nama Lengkap = ", Nama_Lengkap("Dhandi", "Adam"))
}
