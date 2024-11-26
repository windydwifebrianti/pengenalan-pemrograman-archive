package main

import "fmt"

func main() {

	var n, Sisi int

	fmt.Println("Banyaknya persegi: ")
	fmt.Scan(&n)

	fmt.Println("Sisi persegi: ")

	for i := 1; i <= n; i++ {

		fmt.Scan(&Sisi)

		luas := Sisi * Sisi
		keliling := 4 * Sisi
		fmt.Println(luas, keliling)
	}
}
