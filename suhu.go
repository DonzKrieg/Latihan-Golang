package main

import "fmt"

func main() {
	var celAwal, celAkhir, step, c, r, f, k, i float32
	fmt.Scanln(&celAwal, &celAkhir, &step)
	fmt.Printf("%10s %10s %10s %10s\n", "Celcius", "Reamur", "Fahrenheit", "Kelvin")

	for i = 1; i <= celAkhir/step; i++ {
		c = celAwal
		r = reamur(celAwal)
		f = fahrenheit(celAwal)
		k = kelvin(celAwal)
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", c, r, f, k)
		celAwal += step
	}
}

func reamur(c float32) float32 {
	var rea float32
	rea = c * 4 / 5
	return rea
}

func fahrenheit(c float32) float32 {
	var fah float32
	fah = (c * 9 / 5) + 32
	return fah
}

func kelvin(c float32) float32 {
	var kel float32
	kel = c + 273
	return kel
}
