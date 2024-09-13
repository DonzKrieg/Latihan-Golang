package main

import "fmt"

func main() {
	var input int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&input)

	var bilanganPertama int = 0
	var bilanganKedua int = 1

	for i := 0; i <= input; i++ {
		x := bilanganPertama
		bilanganPertama = bilanganKedua
		bilanganKedua = x + bilanganKedua

		fmt.Print(x, " ")
	}
}
