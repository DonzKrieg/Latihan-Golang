package main

import "fmt"

func main() {

	var x, y int
	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)

	y = ((3 * x) - 5) * ((2 * x) + 1)
	fmt.Println("Nilai y adalah: ", y)

}
