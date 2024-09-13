package main

import "fmt"

// Tipe bentukan struct titik dengan field x dan y (float64) dan warna (string)
type titik struct {
	x, y  float64
	warna string
}

func main() {
	var x1, y1, x2, y2 float64
	var w1, w2 string
	var t1, t2 titik

	// Baca data x1, y1, w1, x2, y2, w2
	fmt.Println("Masukkan koordinat dan warna titik 1:")
	fmt.Scan(&x1, &y1, &w1)
	fmt.Println("Masukkan koordinat dan warna titik 2:")
	fmt.Scan(&x2, &y2, &w2)

	// Pembuatan titik t1
	t1 = pembuatanTitikBaru(x1, y1, w1)

	// Pembuatan titik t2
	t2 = pembuatanTitikBaru(x2, y2, w2)

	// Pencetakan titik t1 dan t2
	fmt.Printf("Data titik 1: Koordinat (%.1f, %.1f), warna %s\n", t1.x, t1.y, t1.warna)
	fmt.Printf("Data titik 2: Koordinat (%.1f, %.1f), warna %s\n", t2.x, t2.y, t2.warna)
}

func pembuatanTitikBaru(x, y float64, w string) titik {
	/* Mengembalikan sebuah titik dengan koordinat x dan y, serta warna w */
	return titik{x, y, w}
}
