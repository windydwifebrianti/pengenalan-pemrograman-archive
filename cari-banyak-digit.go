package main

import "fmt"

func main() {
	var angka, i, 
	banyakDigit int
	
	fmt.Scanf("%d", &angka)

	i = angka
	banyakDigit = 0

	for i > 0 {
		i = (i - (i % 10)) / 10
		banyakDigit++

		fmt.Printf("banyak Digit: %d\n", banyakDigit)

	}
}
