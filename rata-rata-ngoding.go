package main

import "fmt"

func main() {

	var hari, jamPerHari, total int
	var rata float64

	fmt.Println("Banyak hari: ")
	fmt.Scan(&hari)

	jamPerHari = 0

	for i := 1; i <= hari; i++ {
		fmt.Scan(&jamPerHari)

		total += jamPerHari
	}
	rata = float64(total) / float64(hari)

	fmt.Println(rata)
}
