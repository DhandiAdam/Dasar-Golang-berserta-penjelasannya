package main

import "fmt"

func main() {
	// Jadi kalau misal kita memakai slice [2:6] maka akan ada output dari 2 sampai 6
	//[] semua itu tergantung dari jumlah yang kamu input sebelumnya

	//Kalau di program itu 0		 1		  2			  3		  4			5
	//Kalau di Manusia itu 1		 2		  3			  4		  5			6
	names := [...]string{"Dhandi", "Adam", "Syaifudin", "Etoy", "otoy", "Nasrullah"}

	values1 := names[2:6]

	fmt.Println(values1)
	Angka := [...]int{1, 2, 3, 4, 5, 6}

	Range := Angka[2:5]

	fmt.Println("Angka berapa yang ke cetak ? ", Range)

	RangeNameFirstandLast := names[3:]
	//Nama yang tergantung kamu input berapa
	fmt.Println("Names yang didepan sampai belakang ", RangeNameFirstandLast)

	RangeNameAll := names[:]
	fmt.Println(RangeNameAll)

	fmt.Println(" ")
	RangeAngkaFirstandLast := Angka[3:]
	//Angka yang tergantung kamu input berapa
	fmt.Println("Angka yang didepan sampai belakang ", RangeAngkaFirstandLast)

	RangeAngkaAll := Angka[:]
	fmt.Println(RangeAngkaAll)
}
