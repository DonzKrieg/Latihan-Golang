package main

import (
	"fmt"
	"math"
)

func main() {
	var r, t, x, y float64
	fmt.Scanln(&r, &t)
	x = panjangX(r, t)
	y = panjangY(r, t)
	fmt.Printf("%.2f %.2f\n", x, y)
}

func panjangX(r, t float64) float64 {
	var x float64
	x = r * math.Cos(math.Pi/180*t)
	return x
}

func panjangY(r, t float64) float64 {
	var y float64
	y = r * math.Sin(math.Pi/180*t)
	return y
}
