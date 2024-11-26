package main

import "fmt"

func main() {

	var jml int

	jml = 0
	for i := 1; i <= 50; i += 2 {
		jml += i
	}
	fmt.Println(jml)
}
