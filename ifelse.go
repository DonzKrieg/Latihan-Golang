package main

import "fmt"

func main() {

	fmt.Print("Masukkan nilaimu disini: ")

	var nilai float64
	fmt.Scan(&nilai)

	if nilai >= 80.5 {

		fmt.Println("Nilai akhirnya adalah A")

	} else if nilai < 80.5 && nilai >= 70 {

		fmt.Println("Nilai akhirnya adalah B")

	} else if nilai < 70 && nilai >= 55.9 {

		fmt.Println("Nilai akhirnya adalah C")

	} else {

		fmt.Println("Tidak Lulus!!!")

	}

}
