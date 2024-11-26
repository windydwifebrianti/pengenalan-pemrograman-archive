package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan n bulan: ")
	fmt.Scanln(&n)
	totalTabungan := 5000

	total := totalTabungan * n
	fmt.Println(total)

	fmt.Println("Total tabungan, adalah", total)

}
