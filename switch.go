package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	fmt.Print("Tulis ", i, " sebagai ")
	switch i {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	case 3:
		fmt.Println("tiga")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Sekarang adalah weekend")
	default:
		fmt.Println("Sekarang adalah weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Sekarang masih belum jam dua belas")
	default:
		fmt.Println("Sekarang sudah lewat jam dua belas")
	}

	switch {
	case uint8(t.Month()) == 1:
		fmt.Println("Sekarang adalah bulan Januari")
	case uint8(t.Month()) == 2:
		fmt.Println("Sekarang adalah bulan Februari")
	case uint8(t.Month()) == 3:
		fmt.Println("Sekarang adalah bulan Maret")
	case uint8(t.Month()) == 4:
		fmt.Println("Sekarang adalah bulan April")
	case uint8(t.Month()) == 5:
		fmt.Println("Sekarang adalah bulan Mei")
	case uint8(t.Month()) == 6:
		fmt.Println("Sekarang adalah bulan Juni")
	case uint8(t.Month()) == 7:
		fmt.Println("Sekarang adalah bulan Juli")
	case uint8(t.Month()) == 8:
		fmt.Println("Sekarang adalah bulan Agustus")
	case uint8(t.Month()) == 9:
		fmt.Println("Sekarang adalah bulan September")
	case uint8(t.Month()) == 10:
		fmt.Println("Sekarang adalah bulan Oktober")
	case uint8(t.Month()) == 11:
		fmt.Println("Sekarang adalah bulan November")
	case uint8(t.Month()) == 12:
		fmt.Println("Sekarang adalah bulan Desember")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Aku adalah boolean")
		case int:
			fmt.Println("Aku adalah integer")
		default:
			fmt.Printf("Tipe variabelnya adalah: %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(12)
	whatAmI("Seharusnya ini adalah string.")
}
