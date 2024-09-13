package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	var valid bool = true

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&input)

	for i := 0; i < len(input)-1; i++ {
		if valid == false {
			break
		}
		prev, _ := strconv.Atoi(input[i : i+1])
		next, _ := strconv.Atoi(input[i+1 : i+2])
		valid = next == prev+1 || next == prev-1
	}
	fmt.Print(valid)
}
