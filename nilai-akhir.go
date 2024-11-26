package main 

import "fmt"

func main () {

	var tugas, kuis, ujian, praktikum float64
	var nilaiAkhir float64

	fmt.Println("Masukkan nilai tugas")
	fmt.Scan(&tugas)

	fmt.Println("Masukkan nilai kuis")
	fmt.Scan(&kuis)

	fmt.Println("Masukkan nilai ujian")
	fmt.Scan(&ujian)

	fmt.Println("Masukkan nilai praktikum")
	fmt.Scan(&praktikum)

	nilaiAkhir = (0.1 * tugas) + (0.2 * kuis) + (0.5 * ujian) + (0.2 * praktikum)

    fmt.Println("Nilai Akhir: ",nilaiAkhir)

}