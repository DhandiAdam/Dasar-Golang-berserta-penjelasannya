package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a int
	var b int

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan Nilai A = ")
	inputA, _ := reader.ReadString('\n')
	inputA = strings.TrimSpace(inputA)

	a, err := strconv.Atoi(inputA)

	if err != nil {
		fmt.Println("Harus diawali hurf ", err)
		return
	}
	fmt.Println("Nilai Anda ", a)
	fmt.Println("")
	fmt.Print("Masukan Nilai B = ")
	inputB, _ := reader.ReadString('\n')
	inputB = strings.TrimSpace(inputB)

	b, err1 := strconv.Atoi(inputB)
	if err1 != nil {
		fmt.Println("Harus diawali Angka ", err1)
		return
	}
	fmt.Println("Nilai Anda Adalah = ", b)
	fmt.Println("")
	var C int
	fmt.Print("Masukan Nilai C = ")
	inputC, _ := reader.ReadString('\n')
	inputC = strings.TrimSpace(inputC)

	C, err2 := strconv.Atoi(inputC)
	if err2 != nil {
		fmt.Println("Harus diawali Angka ", err2)
		return
	}
	fmt.Println("Nilai Anda adalah = ", C)
	fmt.Println("Perbandingan A dan b ")
	var result1 bool
	result1 = a < b
	fmt.Println("Apakah A lebih kecil dari nilai B (Jika iya maka true) ", result1)
	result1 = a == b
	fmt.Println("Apakah Nilai a == b sama? (jika sama maka true) ", result1)
	result1 = a != b
	fmt.Println("Apakah Nilai A tidak sama dengan b a != b  (jika tidak sama dengan maka true) ", result1)
	fmt.Println("=============================")
	result1 = a > C
	fmt.Println("Apakah A lebih besar dari nilai C (Jika iya maka true) ", result1)
	result1 = a < C
	fmt.Println("Apakah A lebih kecil dari nilai C (Jika iya maka true) ", result1)
	result1 = a == C
	fmt.Println("Apakah Nilai a == C sama? (jika sama maka true) ", result1)
	result1 = a != C
	fmt.Println("Apakah Nilai A tidak sama dengan C a != C  (jika tidak sama dengan maka true) ", result1)

}
