package main

import "fmt"

// tipe bentukan struct mobil dengan field merek, tahun_produksi, dan kecepatan
type mobil struct {
	merek         string
	tahunProduksi int
	kecepatan     int
}

func main() {
	var m1, m2, m3 mobil
	var rataRataKecepatan float64

	// Meminta input data untuk 3 mobil
	fmt.Println("Masukkan data untuk mobil 1:")
	fmt.Scan(&m1.merek, &m1.tahunProduksi, &m1.kecepatan)
	fmt.Println("Masukkan data untuk mobil 2:")
	fmt.Scan(&m2.merek, &m2.tahunProduksi, &m2.kecepatan)
	fmt.Println("Masukkan data untuk mobil 3:")
	fmt.Scan(&m3.merek, &m3.tahunProduksi, &m3.kecepatan)

	// hitung rata-rata kecepatan dari 3 mobil
	rataRataKecepatan = float64(m1.kecepatan+m2.kecepatan+m3.kecepatan) / 3

	// cetak data data mobil dengan rata-rata kecepatannya
	fmt.Printf("Rata-rata kecepatan mobil %s (%d), ", m1.merek, m1.tahunProduksi)
	fmt.Printf("mobil %s (%d), dan mobil %s (%d): ", m2.merek, m2.tahunProduksi, m3.merek, m3.tahunProduksi)
	fmt.Printf("%.2f\n", rataRataKecepatan)
}
