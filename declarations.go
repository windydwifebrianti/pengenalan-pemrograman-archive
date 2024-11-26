package main

import "fmt"

func main() {
	type NoHP string

	var nomerwindy NoHP = "0895808120302"

	var contoh string = "081327327259"
	var contohHP NoHP = NoHP(contoh)

	fmt.Println(nomerwindy)
	fmt.Println(contohHP
}
