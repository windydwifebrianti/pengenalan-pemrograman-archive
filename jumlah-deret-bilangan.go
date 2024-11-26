package main

import "fmt"

func main() {
	var n, m int
	var total int = 0

	// Input bilangan n
	fmt.Scan(&n)
	// Input bilangan m
	fmt.Scan(&m)

	for i := n; i <= m; i++ {
		total = total + 2/i
	}

	fmt.Println(total)
}
