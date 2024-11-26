package main

import "fmt"

func main() {

	var n int

	fmt.Println("Masukkan n putaran mengelilingi lapangan")
	fmt.Scan(&n)

	total := 0

	fmt.Println("Masukkan catatan waktu tiap putaran dalam satuan detik")

	for i := 0; i <= n; i++ {
		fmt.Printf("Putaran %d: ", i+1)

		fmt.Scan(&total)
	}

	Rata := 0
	for i := 0; i <= n; i += n {
		Rata = total / n
	}
	fmt.Println("Rata-rata: ", Rata)
}
