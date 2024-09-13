package main

import "fmt"

func main() {
	var item []bool
	var person int
	fmt.Scan(&person)

	for i := 0; i < person; i++ {
		var a, b, c, d, e bool
		fmt.Scan(&a, &b, &c, &d, &e)
		item = append(item, a)
		item = append(item, b)
		item = append(item, c)
		item = append(item, d)
		item = append(item, e)
	}

	var benar bool
	for i := 0; i < len(item)-1; i++ {
		benar = item[i] == true
		if !benar {
			break
		}
	}

	fmt.Print(benar)
}
