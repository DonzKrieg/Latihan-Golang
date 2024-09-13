package main

import "fmt"

const N int = 1000
type tabData [N]int

func main() {
	var p tabData
	var n int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i])
	}
}

func urut(A, n int) bool {
	
}