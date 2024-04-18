package main

import "fmt"

func main() {
	//Kalau di program itu 0		 1		  2		  3		  4			5		  7
	//Kalau di Manusia itu 1		 2		  3		  4		  5			6		  7
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	DaySlice := days[5:]
	fmt.Println(days)
	DaySlice[0] = "Sabtu Baru"
	DaySlice[1] = "Minggu Baru"
	//Jika kalian menggubah data slicenya maka semuanya akan kerubah
	fmt.Println("Data yang berubah ", days)
	fmt.Println(DaySlice) //Akan diganti dengan yang baru

	DaySlice1 := append(DaySlice)
	DaySlice1[0] = "Ups"
	fmt.Println(DaySlice1)

	DaySlice2 := append(DaySlice, "Horeee Liburannn")
	DaySlice2[0] = "Ups"
	fmt.Println(DaySlice2)
}
