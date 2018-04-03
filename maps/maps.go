package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// Pra adalah variable untuk key, dan prs adalah variable untuk valuenya
	pra, prs := m["k2"]
	fmt.Println("pra:", pra)
	fmt.Println("prs:", prs)

	// Kita bisa mengganti variable pra dengan underscore jika tidak ingin memakainya
	// Alasanya karena golang tidak bisa membiarkan satu variabel pun tidak terpakai,
	// sehingga harus ada sesuatu untuk mengganti variabel tersebut
	_, prs2 := m["k2"]
	fmt.Println("prs2:", prs2)

	fmt.Println("m['k1']")
	x, y := m["k1"]
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
