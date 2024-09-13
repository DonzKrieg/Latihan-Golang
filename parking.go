package main

import (
	"fmt"
)

type Time struct {
	Hour   int
	Minute int
	Second int
}

func main() {
	// Input waktu masuk parkir
	var masuk Time
	fmt.Println("Masukkan waktu masuk parkir (jam, menit, detik):")
	fmt.Scan(&masuk.Hour, &masuk.Minute, &masuk.Second)

	// Input waktu keluar parkir
	var keluar Time
	fmt.Println("Masukkan waktu keluar parkir (jam, menit, detik):")
	fmt.Scan(&keluar.Hour, &keluar.Minute, &keluar.Second)

	// Menghitung lama waktu parkir
	lamaParkir := hitungLamaParkir(masuk, keluar)

	// Menampilkan lama waktu parkir
	fmt.Printf("Lama waktu parkir: %d jam, %d menit, %d detik\n", lamaParkir.Hour, lamaParkir.Minute, lamaParkir.Second)
}

func hitungLamaParkir(masuk, keluar Time) Time {
	var lama Time

	// Menghitung detik lama parkir
	detikLamaParkir := (keluar.Hour*3600 + keluar.Minute*60 + keluar.Second) - (masuk.Hour*3600 + masuk.Minute*60 + masuk.Second)

	// Mengkonversi detik menjadi jam, menit, detik
	lama.Hour = detikLamaParkir / 3600
	lama.Minute = (detikLamaParkir % 3600) / 60
	lama.Second = (detikLamaParkir % 3600) % 60

	return lama
}
