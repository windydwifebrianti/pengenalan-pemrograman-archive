package main

import "fmt"

func main() {
	var n int
	var nama string
	fmt.Println("Banyaknya percetakan string: ")
	fmt.Scan(&n)

	fmt.Println("Nama lu saha: ")
	fmt.Scan(&nama)

	for i := 1; i <= n; i++ {
		fmt.Println(nama)
	}
}
