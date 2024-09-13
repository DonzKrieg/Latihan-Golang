package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	var total int = 0

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&input)

	// 1 2 3 4 5
	for i := len(input) - 1; i >= 0; i-- {
		x, _ := strconv.Atoi(input[i : i+1])
		total += x
		fmt.Printf("%d ", x)
	}
	fmt.Print("\n", total)

}
