package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	faktor(1, n)
}

func faktor(i, n int) int {
	if i > n {
		return n
	}
	if n%i == 0 {
		fmt.Printf("%d ", i)
	}
	return faktor(i+1, n)
}
