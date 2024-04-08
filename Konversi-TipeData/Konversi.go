package main

import "fmt"

func main() {
	//Jika ingin mengkonversi sebuah nilai maka harus sesuai kapasitas tipe datanya \
	// Jika dia menampung didalam sebuah tipe data akan menghasilkan nilai
	//Jika dia tidak mampu menampung datanya maka dia akan minus
	var nilai32 int32 = 3278
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name string = "Dhandi"
	//name[0] menandakan kita ingin menggambil sebuah nilai di index 0 yaitu D
	//jadi bisa kita sebut seperti ini D = index 0, h = index 1, a = index 2 dst
	var e = name[0]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println("Indeks[0] maka akan muncul = ", eString)

	// oke disini saya ingin mengimplementasikan
	index1 := name[1]
	var INDEX1 = string(index1)

	indext2 := name[2]
	var INDEKS2 = string(indext2)

	fmt.Println("Indeks[1] maka akan muncul = ", INDEX1)
	fmt.Println("Indeks[2] maka akan muncul = ", INDEKS2)

}
