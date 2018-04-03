package main

import "fmt"

func main() {
	fmt.Println("Menghitung luas dan keliling persegi panjang.")

	var panjang, lebar int

	fmt.Printf("Panjang: ")
	fmt.Scan(&panjang)
	fmt.Printf("Lebar: ")
	fmt.Scan(&lebar)

	luas := large(panjang, lebar)
	keliling := around(panjang, lebar)

	fmt.Printf("\nLuas: %v\nKeliling: %v", luas, keliling)
}

func large(panjang int, lebar int) int {
	return panjang * lebar
}

func around(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}
