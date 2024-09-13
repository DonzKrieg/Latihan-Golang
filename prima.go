package main

import "fmt"

func main() {

	var x, i int
	var prima bool
	fmt.Scan(&x)
	prima = true

	if x == 0 || x == 1 {
		prima = false
	} else {
		for i = 2; i <= x/2; i++ {
			if x%i == 0 {
				prima = false
				break
			}
		}
	}

	fmt.Print(prima)

}
