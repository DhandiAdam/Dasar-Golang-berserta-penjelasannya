package main

import "fmt"

func main() {
	var newSlice []string = make([]string, 3, 5)

	newSlice[0] = "Dhandi"
	newSlice[1] = "Adam"
	newSlice[2] = "Suhudin"

	fmt.Println(newSlice)
	//Jika ingin menambahkan sebuah data make memakai append
	newSlicee2 := append(newSlice, "Dhandi")
	fmt.Println(newSlicee2)
	fmt.Println("Panjang Data yang di input", len(newSlicee2))
	fmt.Println("Capacity dari sebuah data ", cap(newSlicee2))

}
