package main

import "fmt"

func main() {
   
    var a int
    var b int
    var hasil int

	fmt.Println("Masukkan nilai a")
    fmt.Scan(&a)

    b = 2 * a

    hasil = b + 10

    fmt.Println("Hasil\n", hasil)
}