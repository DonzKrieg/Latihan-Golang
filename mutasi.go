package main

import "fmt"

func main() {

	var x, y int
	fmt.Scanln(&x, &y)
	fmt.Println(facX(x), facY(y), mut(x, y))

}

func facX(p int) int {
	var i, fac int
	fac = 1
	for i = p; i >= 1; i-- {
		fac *= i
	}
	return fac
}

func facY(q int) int {
	var i, fac int
	fac = 1
	for i = q; i >= 1; i-- {
		fac *= i
	}
	return fac
}

func mut(p, q int) int {
	var i, fac, mutasi int
	fac = 1
	if p >= q {
		for i = p - q; i >= 1; i-- {
			fac *= i
		}
		mutasi = facX(p) / fac
	} else {
		for i = q - p; i >= 1; i-- {
			fac *= i
		}
		mutasi = facY(q) / fac
	}

	return mutasi
}
