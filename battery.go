package main

import "fmt"

func main() {

	var E0, E1, c, cc int
	fmt.Scan(&E0)
	fmt.Scan(&c)
	fmt.Scan(&E1)

	cc = (E0 - E1) / c
	fmt.Println(cc)

}
