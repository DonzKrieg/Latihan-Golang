package main

import "fmt"

func main() {

	var totalCangkir int = 0

	var persediaanGula, persediaanKopi int
	var butuhGula, butuhKopi int

	fmt.Scanln(&persediaanGula, &persediaanKopi, &butuhGula, &butuhKopi)

	for butuhGula <= persediaanGula && butuhKopi <= persediaanKopi {
		persediaanGula -= butuhGula
		persediaanKopi -= butuhKopi
		totalCangkir++
	}
	fmt.Print(totalCangkir)

}
