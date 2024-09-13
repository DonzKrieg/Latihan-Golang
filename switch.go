package main

import "fmt"

func main() {

	var n, name string
	fmt.Scan(&n)

	name = n

	switch name {

	case "ihsan":
		fmt.Println("Hello Ihsan!")

	case "joko":
		fmt.Println("Hello joko!")

	default:
		fmt.Println("May i know u?")

	}

	// short statement
	switch length := len(name); length > 5 {

	case true:
		fmt.Println("Nama terlalu panjang")

	case false:
		fmt.Println("Nama sudah benar")

	}

}
