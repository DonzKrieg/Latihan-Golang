package main

import "fmt"

func pecahUang(uang int, k10, k2, k1 *int) {
	*k10 = uang / 10000
	*k2 = uang % 10000 / 2000
	*k1 = uang % 10000 % 2000 / 1000
}

func main() {
	var duit, d10, d2, d1 int
	fmt.Scan(&duit)
	pecahUang(duit, &d10, &d2, &d1)
	fmt.Println(d10, "Lembar 10.000")
	fmt.Println(d2, "Lembar 2.000")
	fmt.Println(d1, "Lembar 1.000")
}
