package main

import "fmt"

func main() {

	var r, pi, bola float64
	fmt.Print("Masukkan jari jari lingkaran: ")
	fmt.Scan(&r)

	pi = 3.14
	bola = 4 / 3 * pi * r * r * r
	fmt.Println("Volume bola adalah:", bola)

}
