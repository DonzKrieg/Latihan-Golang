package main

import "fmt"

func main() {

	var x int
	var bil bool
	fmt.Scan(&x)

	bil = genap(x)
	fmt.Println(bil)
}

func genap(p int) bool {
	var bilGen bool
	if p%2 == 0 {
		bilGen = true
	} else {
		bilGen = false
	}
	return bilGen
}
