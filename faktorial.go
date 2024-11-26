package main

import "fmt"

func main() {

	var n int

	fmt.Println("Masukkan bilangan bulat positif")
	fmt.Scan(&n)

	total := 1

	for i := 0; i <= n-2; i++ {

		total = total * (n - i)

	}
	fmt.Println(total)
}
