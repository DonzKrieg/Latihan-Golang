package main

import "fmt"

func main() {

	var stel, neb, gal, qua, lum int

	fmt.Scan(&lum)

	qua = lum / 1200
	gal = lum % 1200 / 120
	neb = lum % 120 / 40
	stel = lum % 40 / 20
	lum = lum % 20

	fmt.Println(qua, gal, neb, stel, lum)

}
