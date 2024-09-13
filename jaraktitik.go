package main

import (
	"fmt"
	"math"
)

// Definisikan struktur Titik
type Titik struct {
	x float64
	y float64
}

// Fungsi untuk menghitung jarak antara dua titik
func jarak(P1, P2 Titik) float64 {
	a := P2.x - P1.x
	b := P2.y - P1.y
	return akar(a*a + b*b)
}

// Fungsi untuk menghitung akar kuadrat
func akar(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	// Input
	var x1, y1, x2, y2 float64
	fmt.Println("Masukkan nilai x dan y untuk titik P1:")
	fmt.Print("x: ")
	fmt.Scan(&x1)
	fmt.Print("y: ")
	fmt.Scan(&y1)
	fmt.Println("Masukkan nilai x dan y untuk titik P2:")
	fmt.Print("x: ")
	fmt.Scan(&x2)
	fmt.Print("y: ")
	fmt.Scan(&y2)

	// Membuat objek dari struktur Titik
	P1 := Titik{x1, y1}
	P2 := Titik{x2, y2}

	// Memanggil fungsi jarak
	hasilJarak := jarak(P1, P2)

	// Output
	fmt.Println("Jarak antara dua titik adalah:", hasilJarak)
}
