package main

import "fmt"

func main() {

	var N int

	fmt.Println("Masukkan bilangan positif ")
	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		if N%i == 0 {
			fmt.Printf("%d true\n", i)
		} else {
			fmt.Printf("%d false\n", i)
		}
	}

}
