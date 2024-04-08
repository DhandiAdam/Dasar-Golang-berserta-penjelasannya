package main

import "fmt"

func main() {
	//Disini kita melakukan function terhadap operator matematika
	Operasi_Penjumlahan := func(A, B int) int {
		result := A + B
		return result
	}

	Operasi_Pengurangan := func(C, D int) int {
		pengurangan := Operasi_Penjumlahan(10, 20)
		result := pengurangan - C - D
		return result
	}

	Operasi_Perkalian := func(E, F int) int {
		Perjumlahan := Operasi_Penjumlahan(10, 20)
		Pengurangan := Operasi_Pengurangan(5, 5)
		result := (Perjumlahan - Pengurangan) * (E * F)
		return result
	}

	Operasi_Pembagian := func(A int) int {
		Perjumlahan := Operasi_Penjumlahan(10, 20)
		Pengurangan := Operasi_Pengurangan(5, 5)
		Perkalian := Operasi_Perkalian(5, 5)
		RESULT := (Perjumlahan - Pengurangan) * (Perkalian) / A
		return RESULT
	}

	fmt.Println("10 + 20 = ", Operasi_Penjumlahan(10, 20))
	fmt.Println("30 - 5 - 5 = ", Operasi_Pengurangan(5, 5))
	fmt.Println("Hasil Dari (30-20)*5*5 = ", Operasi_Perkalian(5, 5))
	fmt.Println("Hasil Dari (30-20)*(250) / 2 = ", Operasi_Pembagian(2))

}
