package main

import "fmt"

func main() {

	var p, q, r bool

	p = false
	q = false
	r = true

	prop := p && q || r
	fmt.Println("Pernyataannya adalah", prop)

}
