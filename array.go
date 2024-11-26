package main

import "fmt"

func main() {

	lahan := [5]int{
		50,
		60,
		55,
		70,
		65,
	}

	jumlahTotal := 0

	for i := 0; i < len(lahan); i++ {
		jumlahTotal += lahan[i]

	}

	fmt.Println(jumlahTotal)

	max := 0

	for i := 0; i < len(lahan); i++ {

		if lahan[i] > max {
			max = lahan[i]
		}
	}
	fmt.Println(max)

	rata := 0
	rata = jumlahTotal / len(lahan)

	fmt.Println(rata)

}
