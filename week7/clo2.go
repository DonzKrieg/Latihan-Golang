package main

import (
	"fmt"
	"strconv"
)

func main() {

	var code string
	var disc, cBack bool
	fmt.Print("Masukkan kode kupon anda: ")
	fmt.Scan(&code)

	bilTeng, _ := strconv.Atoi(code[1:3])
	bil1, _ := strconv.Atoi(code[0:1])
	bil2, _ := strconv.Atoi(code[3:4])
	bil4, _ := strconv.Atoi(code[3:4])

	disc = bilTeng%2 == 0

	if bil4 != 0 {

		cBack = (bil1+bil2)%bil4 == 0

	}

	fmt.Println("Discount? ", disc)
	fmt.Println("Cashback? ", cBack)

}
