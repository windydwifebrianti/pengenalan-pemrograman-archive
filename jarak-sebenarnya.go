package main

import "fmt"

func main() {
    
    var jarak_peta int = 8
    const skala int = 300000
    var jarak_sebenarnya float64

	jarak_sebenarnya = float64(jarak_peta * skala) / 100000
    
    fmt.Printf("%.2f\n", jarak_sebenarnya)
}