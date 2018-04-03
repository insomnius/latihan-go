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

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
