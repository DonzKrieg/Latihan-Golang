package main

import "fmt"

func main() {

	var belanja, diskon float32
	var member bool
	fmt.Scanln(&belanja, &member)

	diskon = total(belanja, member)
	fmt.Println(diskon)

}

func total(x float32, y bool) float32 {
	var hasil float32
	if x > 100000 && y == true {
		hasil = x - (x * 0.1)
	} else if x > 100000 && y == false {
		hasil = x - (x * 0.05)
	} else {
		hasil = x
	}
	return hasil
}
