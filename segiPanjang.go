package main

import "fmt"

type Reactangle struct {
	length, width int
	color         string
	property      Geometry
}

type Geometry struct {
	Luas     int
	Keliling int
}

func isidata(persegi *Reactangle) {
	fmt.Printf("Length: %d, Width: %d, Color: %s \n", persegi.length, persegi.width, persegi.color)
}

func hitung(persegi *Reactangle) {
	persegi.property.Luas = persegi.length * persegi.width
	persegi.property.Keliling = 2 * (persegi.length + persegi.width)
}

func main() {
	var persegi Reactangle

	fmt.Scan(&persegi.length, &persegi.width, &persegi.color)

	isidata(&persegi)
	hitung(&persegi)

	fmt.Println("Hasil Luas: ", persegi.property.Luas)
	fmt.Println("Hasil Keliling: ", persegi.property.Keliling)

}
