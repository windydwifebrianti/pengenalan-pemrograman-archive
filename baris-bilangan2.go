package main

import "fmt"

func main() {

	var n, m int

	fmt.Println("Masukkan nilai n (awal)")
	fmt.Scan(&n)

	fmt.Println("Masukkan nilai m (batas)")
	fmt.Scan(&m)

	fmt.Println("Barisan Bilangan")

	for i := n; i < m; i++ {
		fmt.Println(i)
	}
}
