package main

import "fmt"

func converted(c float32, r, f, k *float32) {
	*r = c * 4 / 5
	*f = (c * 9 / 5) + 32
	*k = c + 273.15
}

func main() {
	var cel, rea, fah, kel float32
	fmt.Scan(&cel)
	converted(cel, &rea, &fah, &kel)
	fmt.Printf("%.2f R %.2f F %.2f K\n", rea, fah, kel)
}
