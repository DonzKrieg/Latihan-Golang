package main

import "fmt"

func main() {
	var p, l, k int
	fmt.Scan(&p)
	fmt.Scan(&l)

	k = 2 * (p + l)
	fmt.Println(k)
}
