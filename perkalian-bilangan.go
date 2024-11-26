package main

import "fmt"

func main() {
    var n int
    var hasil int = 1  
    var bilangan int  
    
	fmt.Println("Masukkan banyaknya bilangan bulat")
    fmt.Scan(&n)
   
    for i := 0; i < n; i++ {

		fmt.Println("Masukkan bilangan bulat")
        fmt.Scan(&bilangan)

        hasil *= bilangan
    }

    fmt.Println("Hasil", hasil)
}