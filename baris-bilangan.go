package main

import "fmt"

func main() {

	var n, m uint16

	fmt.Println("Lebokno nilai n (awal)")
	fmt.Scan(&n)

	fmt.Println("Lebokno nilai m (batas)")
	fmt.Scan(&m) 

	fmt.Println("Barisan bilangan") 
	for n<m {
		fmt.Println(n)
		n++
	}
}
