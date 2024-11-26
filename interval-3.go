package main

import "fmt"

func main() {

	var n int

	fmt.Print("Masukkan bilangan bulat n:")

	fmt.Scanln(&n)

	for i := 3; i <= n; i += 3 {

		fmt.Println(i)

	}

}
