package main

import (
	"fmt"
)

func main() {

	if 7%2 == 0 {
		fmt.Println("Angka 7 adalah genap.")
	} else {
		fmt.Println("Angka 7 adalah ganjil.")
	}

	if 8%4 == 0 {
		fmt.Println("8 bisa dibagi habis oleh 4")
	} else {
		fmt.Println("8 tidak bisa dibagi habis oleh 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "Negatif")
	} else if num < 10 {
		fmt.Println(num, "Lebih dari satu digit")
	} else {
		fmt.Println(num, "Banyak digit")
	}

	var nilai uint8 = 90

	if nilai < 40 {
		fmt.Println("Nilai Anda adalah:", nilai, ", grade yang Anda dapat adalah: 'E'")
	} else if nilai < 60 {
		fmt.Println("Nilai Anda adalah:", nilai, ", grade yang Anda dapat adalah: 'D'")
	} else if nilai <= 75 {
		fmt.Println("Nilai Anda adalah:", nilai, ", grade yang Anda dapat adalah: 'C'")
	} else if nilai <= 85 {
		fmt.Println("Nilai Anda adalah:", nilai, ", grade yang Anda dapat adalah: 'B'")
	} else if nilai <= 100 {
		fmt.Println("Nilai Anda adalah:", nilai, ", grade yang Anda dapat adalah: 'A'")
	} else {
		var nilai string = "undefined"
		fmt.Println("Nilai Anda adalah:", nilai, ", grade yang Anda dapat adalah: 'undefined'")
	}

}
