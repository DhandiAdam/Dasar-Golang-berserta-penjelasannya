package main

import "fmt"

func main() {
	var values1 = [...]int{
		90, 80,
		80, 90,
		70, 70,
	}
	//[...] Jika Array bisa menginput tanpa ada batasan maksimal
	//[2] jika array ingin menginput 2 kata atau 2 angka

	var values2 = [...]string{
		"Dhandi",
		"Adam",
	}

	fmt.Println(values1)
	fmt.Println(values2)
	fmt.Println("Panjang barisnya values1 = ", len(values1))
	fmt.Println("Panjang barisnya values2 = ", len(values2))

	values1[0] = 100
	fmt.Println(values1)

}
