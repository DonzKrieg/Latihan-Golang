package main

import "fmt"

func main() {

	var (
		a = 20
		b = 2
	)

	z := a % b
	fmt.Println("Hasil: ", z)

	z += 10
	fmt.Println(z)

	z++
	fmt.Println(z)

}
