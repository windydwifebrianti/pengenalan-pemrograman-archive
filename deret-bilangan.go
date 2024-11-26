package main

import "fmt"

func main() {

    var n, m int
    var jumlah int = 0
    
    fmt.Println("Masukkan bilangan n(awal)")
    fmt.Scan(&n)

	fmt.Println("Masukkan bilangan m(batas)")
    fmt.Scan(&m)
    
    if n >= m {
       
        return
    }

	i := n
    for {

        jumlah += i
        i++
        
        if i > m {
            break
        }
    }
    
    fmt.Println("Jumlah deret bilangan dari n-m\n", jumlah)
}