package main

import "fmt"

type HasilGanjil struct {
	BilanganTerakhir int
	JumlahTotal      int
	RataRata         float64
}

func jumlahBilanganGanjil(n int) HasilGanjil {
	var bilangan, jumlah int
	for i := 1; n > 0; i += 2 {
		bilangan = i
		jumlah += i
		n--
	}
	hg := HasilGanjil{
		BilanganTerakhir: bilangan,
		JumlahTotal:      jumlah,
	}
	hg.RataRata = float64(jumlah) / float64(bilangan/2+1)
	return hg
}

func main() {
	var n int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	hasil := jumlahBilanganGanjil(n)

	fmt.Printf("%d %d %.0f\n", hasil.BilanganTerakhir, hasil.JumlahTotal, hasil.RataRata)
}
