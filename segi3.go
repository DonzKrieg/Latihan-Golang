package main

import "fmt"

func main() {

	var sisi1, sisi2, sisi3 int
	var result string
	fmt.Scanln(&sisi1, &sisi2, &sisi3)

	result = cek(sisi1, sisi2, sisi3)
	fmt.Println(result)

}

func cek(x, y, z int) string {
	var hasil string
	if x+y > z && x+z > y && y+z > x {
		hasil = "Segitiga"
	} else {
		hasil = "Bukan Segitiga"
	}
	return hasil
}
