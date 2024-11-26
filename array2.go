package main

import "fmt"

func main() {

	suhu := [10]int{
		33,
		32,
		31,
		30,
		35,
		36,
		35,
		38,
		40,
		39,
	}

	jumlahTotal := 0
	for i := 0; i < len(suhu); i++ {

		jumlahTotal += suhu[i]
	}

	var Rata float64
	Rata = 0
	Rata = float64(jumlahTotal) / float64(len(suhu))

	fmt.Println(Rata)

	suhuTerendah := suhu[0]
	suhuTertinggi := 0
	hariSuhuTerendah := 0
	hariSuhuTertinggi := 0

	for i := 0; i < len(suhu); i++ {

		if suhu[i] < suhuTerendah {
			suhuTerendah = suhu[i]
			hariSuhuTerendah = i+1
		}
	}
	fmt.Println(suhuTerendah)
	fmt.Println(hariSuhuTerendah)

	for i := 0; i < len(suhu);i++ {

		if suhu[i] > suhuTertinggi {
			suhuTertinggi = suhu[i]
			hariSuhuTertinggi = i+1
		}
	}
	fmt.Println(suhuTertinggi)
	fmt.Println(hariSuhuTertinggi)




}
