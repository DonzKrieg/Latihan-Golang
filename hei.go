package main

import "fmt"

func main() {

	var count, i int
	var word string
	fmt.Scan(&count)
	fmt.Scan(&word)

	for i = 1; i <= count; i++ {
		fmt.Println(word)
	}

}
