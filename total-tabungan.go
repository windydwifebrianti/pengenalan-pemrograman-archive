package main

import "fmt"

func main() {

	var n int

	fmt.Print("Masukkan jumlah bulan: ")

	fmt.Scan(&n)

	tpb := 5000

	jmlh := n * tpb

	fmt.Printf("Total tabungan selama %d bulan adalah Rp %d\n", n, jmlh)

}
