package main

import "fmt"

func main() {
	var n int
	var Are string
	fmt.Println("seisun o nyuryoku shite kudasai : ")
	fmt.Scan(&n)

	fmt.Println("input ur name ya anjay: ")
	fmt.Scan(&Are)
	{

		for i := 1; i <= n; i++ {
			fmt.Println(Are)
		}
	}
}
