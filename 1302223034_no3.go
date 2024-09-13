package main

import "fmt"

func main() {
	var huruf byte
	var konsonan bool
	fmt.Scanf("%c", &huruf)
	konsonan = hurufKonsonan(huruf)
	fmt.Println(konsonan)
}

func hurufKonsonan(x byte) bool {
	var up, low bool
	up = (x == 'A' || x == 'I' || x == 'U' || x == 'E' || x == 'O')
	low = (x == 'a' || x == 'i' || x == 'u' || x == 'e' || x == 'o')
	if (x >= 'a' && x <= 'z' && !low) || (x >= 'A' && x <= 'Z' && !up) {
		return true
	} else {
		return false
	}
}
