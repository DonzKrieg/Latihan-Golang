package main

import "fmt"

func main() {

	fmt.Println("Deklarasikan nilai a, b, dan c!!!")

	var a, b, c int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)

	var truth bool = a < b && b > c
	fmt.Print("Kebenaran: ", truth)

}
