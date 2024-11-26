package main

import "fmt"

func main() {

	var hari int
	var jamPerHari int
	var rata float64

	fmt.Println("Masukkan banyak hari")
	fmt.Scan(&hari)

	total := 0
	for i := 1; i <= hari; i++ {

		fmt.Printf("Masukkan banyaknya jam per hari-%d\n", i)
		fmt.Scan(&jamPerHari)

		total += jamPerHari
	}

	rata = float64(total) / float64(hari)
	fmt.Println(rata)
}
