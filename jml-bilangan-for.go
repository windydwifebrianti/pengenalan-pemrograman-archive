package main

import "fmt"

func main() {

	var i, total int

	total = 0
	for i = 1; i <= 50; i++ {
		total += i
	}
	fmt.Println(total)
}