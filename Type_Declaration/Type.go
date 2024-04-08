package main

import "fmt"

func main() {
	type Noktp string

	var ktpDhandi Noktp = "13525433536"

	var contohktp string = "222222222"
	var contoh Noktp = Noktp(contohktp)

	fmt.Println(ktpDhandi)
	fmt.Println(contoh)

}
