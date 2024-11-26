package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Print("Masukkan jumlah putaran: ")

	fmt.Scan(&n)

	waktu := make([]int, n) // Array untuk menyimpan waktu tiap putaran

	totalWaktu := 0

	fmt.Println("Masukkan catatan waktu tiap putaran dalam detik:")

	for i := 0; i < n; i++ {

		fmt.Printf("Putaran %d: ", i+1)

		fmt.Scan(&waktu[i])

		totalWaktu += waktu[i]

	}

	rataRata := float64(totalWaktu) / float64(n)

	fmt.Println("Total waktu berlari:", totalWaktu, "detik")

	fmt.Printf("Rata-rata waktu per putaran: %.2f detik\n", rataRata)

}
