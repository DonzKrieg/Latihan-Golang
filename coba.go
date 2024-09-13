package main

import "fmt"

func main() {

	var l, e, t int
	fmt.Scan(&l, &e)

	t = 0

	for t <= l {

		fmt.Println(false)
		t += e
		fmt.Scan(&e)

	}

	fmt.Println(true)

}
