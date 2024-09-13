package main

import "fmt"

func main() {
	var x byte
	fmt.Scanf("%c", &x)
	fmt.Printf("%c\n", lowToUpper(x))
}

func lowToUpper(kar byte) byte {
	var up byte
	if kar >= 'a' && kar <= 'z' {
		up = kar - 32
	}
	return up
}
