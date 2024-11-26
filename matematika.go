package main

import "fmt"

func main() {
	var a = 5
	var b = 5
	var c = 2
	var d = 3
	var e = a + b - d*c

	fmt.Println(e)

	var i = 10
	i *= 10
	fmt.Println(i)

	i += 5
	fmt.Println(i)
}
