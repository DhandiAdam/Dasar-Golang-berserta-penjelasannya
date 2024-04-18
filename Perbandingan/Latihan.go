package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PilihanProgram menyimpan pilihan program beserta inputannya
type PilihanProgram struct {
	Program string
	Input   string
	Hasil   string
}

func PilihLanjut() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Apakah Anda Ingin Melanjutkan Program (yes/no)? ")
	pilihan, _ := reader.ReadString('\n')
	return strings.TrimSpace(pilihan)
}

func main() {
	history := make([]PilihanProgram, 0) // Slice untuk menyimpan history pilihan pengguna
	for {
		fmt.Println("1. Hitung Luas Segitiga")
		fmt.Println("2. Hitung Persegi")
		fmt.Println("3. Melihat History dan end")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Memilih : ")
		Memilih, _ := reader.ReadString('\n')
		Memilih = strings.TrimSpace(Memilih)

		switch Memilih {
		case "1":
			fmt.Println("Program Luas Segitiga")
			history = append(history, JalankanLuasSegitiga())
		case "2":
			fmt.Println("Program Persegi")
			history = append(history, JalankanPersegi())
		case "3":
			fmt.Println("Program Berakhir")
			fmt.Println("History Pilihan:")
			for i, pilihan := range history {
				fmt.Printf("%d. Program: %s, Input: %s, Hasil: %s\n", i+1, pilihan.Program, pilihan.Input, pilihan.Hasil)
			}
			os.Exit(0)
		default:
			fmt.Println("Program Berakhir")
		}

		pilihan := PilihLanjut()
		if strings.ToLower(pilihan) == "no" {
			fmt.Println("Program Akan Berhenti")
			break
		} else if strings.ToLower(pilihan) != "yes" {
			fmt.Println("Masukan Tidak Valid")
			pilihan2 := PilihLanjut()
			fmt.Println("Apakah anda ingin melanjutkan program ", pilihan2)
		}
	}
}

func HitungLuasSegitiga(inputAlas, TinggiAlas float64) float64 {
	return 0.5 * inputAlas * TinggiAlas
}

func JalankanLuasSegitiga() PilihanProgram {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan Alas Segitiga : ")
	inputAlas, _ := reader.ReadString('\n')
	inputAlas = strings.TrimSpace(inputAlas)

	LuasAlas, err := strconv.ParseFloat(inputAlas, 64)
	if err != nil {
		fmt.Print("Harus Angka Dan gk boleh Huruf ", err)
		return PilihanProgram{Program: "Luas Segitiga", Input: "invalid", Hasil: "invalid"}
	}

	fmt.Print("Masukan Tinggi Segitiga : ")
	TinggiAlas, _ := reader.ReadString('\n')
	TinggiAlas = strings.TrimSpace(TinggiAlas)

	LuasTinggi, err := strconv.ParseFloat(TinggiAlas, 64)
	if err != nil {
		fmt.Print("Harus Angka Dan gk boleh Huruf ", err)
		return PilihanProgram{Program: "Luas Segitiga", Input: "invalid", Hasil: "invalid"}
	}

	Luas := HitungLuasSegitiga(LuasAlas, LuasTinggi)
	fmt.Println("Luas Dari Segitiga ", Luas, "cm²")

	return PilihanProgram{
		Program: "Luas Segitiga",
		Input:   fmt.Sprintf("Alas: %s, Tinggi: %s", inputAlas, TinggiAlas),
		Hasil:   fmt.Sprintf("%.2f cm²", Luas),
	}
}

func HitunglUASPersegi(SisiKiri, SisiKanan int) int {
	return SisiKiri * SisiKanan
}

func JalankanPersegi() PilihanProgram {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan Sisi Kiri : ")
	InputSisiKiri, _ := reader.ReadString('\n')
	InputSisiKiri = strings.TrimSpace(InputSisiKiri)

	SisiKiri, err := strconv.Atoi(InputSisiKiri)
	if err != nil {
		fmt.Print("Harus Angka dan tidak boleh menggunakan Huruf ", err)
		return PilihanProgram{Program: "Persegi", Input: "invalid", Hasil: "invalid"}
	}

	fmt.Print("Masukan SisiKanan : ")
	InputSisiKanan, _ := reader.ReadString('\n')
	InputSisiKanan = strings.TrimSpace(InputSisiKanan)

	SisiKanan, err := strconv.Atoi(InputSisiKanan)
	if err != nil {
		fmt.Print("Harus Angka dan tidak boleh menggunakan Huruf ", err)
		return PilihanProgram{Program: "Persegi", Input: "invalid", Hasil: "invalid"}
	}

	LuasPersegi := HitunglUASPersegi(SisiKiri, SisiKanan)
	KelilingPersegi := 4 * SisiKiri

	fmt.Println("Luas Persegi Adalah : ", LuasPersegi, "cm²")
	fmt.Println("Keliling Persegi Adalah : ", KelilingPersegi, "cm")

	return PilihanProgram{
		Program: "Persegi",
		Input:   fmt.Sprintf("Sisi Kiri: %s, Sisi Kanan: %s", InputSisiKiri, InputSisiKanan),
		Hasil:   fmt.Sprintf("Luas: %d cm², Keliling: %d cm", LuasPersegi, KelilingPersegi),
	}
}
