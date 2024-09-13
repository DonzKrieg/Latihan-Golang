package main

import "fmt"

func main() {

	var letter string
	fmt.Scan(&letter)

	if letter == "A" {
		fmt.Println(5)
	} else if letter == "B" {
		fmt.Println(4)
	} else if letter == "C" {
		fmt.Println(3)
	} else if letter == "D" {
		fmt.Println(2)
	} else if letter == "E" {
		fmt.Println(1)
	} else if letter == "T" {
		fmt.Println(0)
	}

}
