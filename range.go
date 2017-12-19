package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	// Range = Foreach di PHP
	for key, value := range nums {
		sum += value
		fmt.Println("Array ke:", key, "Value:", value)
	}

	fmt.Println("jumlah value:", sum)

}
