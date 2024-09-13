package main

import "fmt"

func main() {

	var x, i int
	fmt.Scan(&x)

	for i = 1; i <= x; i++ {

		if x%i == 0 {
			fmt.Print(" ", i)
		}

	}

}
