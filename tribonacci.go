package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Print(tripplenacci(N))
}

func tripplenacci(N int) int {
	if N <= 3 {
		return 1
	}
	return tripplenacci(N-3) + tripplenacci(N-2) + tripplenacci(N-1)
}