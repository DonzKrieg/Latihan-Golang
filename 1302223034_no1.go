package main

import "fmt"

func main() {
	var nilaiAkhir int
	fmt.Scan(&nilaiAkhir)
	fmt.Println(isValid(nilaiAkhir))
}

func isValid(x int) bool {
	var p bool
	if x < 0 || x > 100 {
		p = false
	} else {
		p = true
	}
	return p
}
