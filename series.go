package main

import (
	"fmt"
	"math"
)

type SeriesReview struct {
	NamaSeries       string
	JumlahEpisode    int
	DurasiPerEpisode int // dalam menit
	DurasiTotal      int // dalam menit
}

func main() {
	var series SeriesReview
	var nJamPerHari int

	// Input data series
	fmt.Println("Masukkan data series (nama, jumlah episode, durasi per episode):")
	fmt.Scan(&series.NamaSeries, &series.JumlahEpisode, &series.DurasiPerEpisode)

	// Menghitung durasi total series dalam menit
	series.DurasiTotal = series.JumlahEpisode * series.DurasiPerEpisode

	// Input jumlah jam yang dapat ditonton dalam sehari
	fmt.Println("Masukkan jumlah jam yang dapat ditonton dalam sehari:")
	fmt.Scan(&nJamPerHari)

	// Menghitung jumlah hari untuk menyelesaikan series
	jumlahHari := hitungJumlahHari(series, nJamPerHari)

	// Menampilkan hasil
	fmt.Printf("Series %s selesai di-review dalam %d hari\n", series.NamaSeries, jumlahHari)
}

func hitungJumlahHari(series SeriesReview, nJamPerHari int) int {
	jumlahJam := float64(series.DurasiTotal) / 60 // Konversi durasi total dari menit ke jam
	return int(math.Ceil(jumlahJam / float64(nJamPerHari)))
}
