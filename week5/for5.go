package main

import "fmt"

func main() {

	var n, i int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {

		var bil bool = n%i == 0
		fmt.Println("Hasil:", i, bil)

	}

}
