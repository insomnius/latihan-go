package main

import "fmt"

// Mengurutkan array dalam satu kali perulangan
func main() {
	fmt.Println("Bubble Short")

	var arrayNumber [20]int

	for i := 0; i < len(arrayNumber); i++ {
		fmt.Printf("Masukan Angka: ")
		fmt.Scan(&arrayNumber[i])
	}

	fmt.Println(arrayNumber)

	i := 0
	for i < len(arrayNumber) {

		if i != (len(arrayNumber) - 1) {
			if arrayNumber[i] > arrayNumber[i+1] {
				x 	:= arrayNumber[i]
				arrayNumber[i] 		= arrayNumber[i+1]
				arrayNumber[i+1]  	= x
				i--
			}
		}

		if i > 0 {
			if arrayNumber[i] < arrayNumber[i-1] {
				x 	:= arrayNumber[i]
				arrayNumber[i] 		= arrayNumber[i-1]
				arrayNumber[i-1]  	= x
				i--
			}
		}

		i++
	}

	fmt.Println("Setelah dilakukan pengurutan.")
	fmt.Println(arrayNumber)
}
