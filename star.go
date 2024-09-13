package main

import "fmt"

func main() {

	var x, i, j int
	fmt.Scan(&x)

	for i = 0; i < x; i++ {
		for j = 0; j < x; j++ {
			if i == j {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Printf("\n")

	}

}
