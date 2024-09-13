package main

import "fmt"

func main() {

	var jari, tinggi, luas float32
	fmt.Scanln(&jari, &tinggi)

	luas = 2*luasLingkaran(jari) + luasSelimut(jari, tinggi)

	fmt.Println(luas)

}

func luasLingkaran(r float32) float32 {
	return 3.14 * (r * r)
}

func luasSelimut(r, t float32) float32 {
	return 2 * 3.14 * r * t
}
