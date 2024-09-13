package main

import "fmt"

func main() {

	var jari, tinggi, vol float32
	fmt.Scanln(&jari, &tinggi)

	vol = vTabung(jari, tinggi)
	fmt.Println(vol)

}

func vTabung(r, t float32) float32 {
	var hasil float32
	hasil = 3.14 * (r * r) * t
	return hasil
}
