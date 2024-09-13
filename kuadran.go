package main

import "fmt"

func main() {

	var x, y float32
	fmt.Scanln(&x, &y)

	if x > 0 && y > 0 {
		fmt.Println("Kuadran I")
	} else if x < 0 && y > 0 {
		fmt.Println("Kuadran II")
	} else if x > 0 && y < 0 {
		fmt.Println("Kuadran IV")
	} else {
		fmt.Println("Kuadran III")
	}

}
