package main

import "fmt"

func main() {
    
    var r int
    const phi float64 = 3.14
    var luas, keliling float64

    fmt.Print("Masukkan jari-jari lingkaran: ")
    fmt.Scanf("%d", &r)

    keliling = 2 * phi * float64(r)

    luas = phi * float64(r) * float64(r)

    fmt.Printf("Keliling Lingkaran: %.2f\n", keliling)
    fmt.Printf("Luas Lingkaran: %.2f\n", luas)
}