package main

import "fmt"

func main() {

	var names = [3]string{

		"Muhammad",
		"Ihsan",
		"Romdhon",
	}

	fmt.Println(names)
	fmt.Println(len(names))

	var codePin [5]int

	codePin[0] = 5
	codePin[1] = 3
	codePin[2] = 1
	codePin[3] = 7
	codePin[4] = 0

	fmt.Println(codePin)
	fmt.Println(len(codePin))

}
