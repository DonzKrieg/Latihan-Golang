package main

import "fmt"

func main() {
	var p1, p2, p3, p4, p5 string
	fmt.Scan(&p1, &p2, &p3, &p4, &p5)
	fmt.Println(manager(p1, p2, p3, p4, p5))
}

func manager(a, b, c, d, e string) string {
	var hasil string
	if a == "kalah" && b == "kalah" && c == "kalah" && d == "kalah" && e == "kalah" {
		hasil = "Dipecat"
	} else {
		hasil = "Tidak Dipecat"
	}
	return hasil
}
