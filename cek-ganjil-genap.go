package main 

import "fmt"

func main() {

	for i := 1; i <= 50; i++ {

		if i%2 == 1 {
			fmt.Printf("%d FALSE\n", i)
		} else {
			fmt.Printf("%d TRUE\n", i)
		}
	}
}