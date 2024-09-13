//Program Permutasi
//kamus
// 	x, y : Integer
//algoritma
// 	input(x, y)
//	output(facX(x), facY(y), mut(x, y))
//endprogram

//function facX(p : Integer) -> Integer
//kamus
//	i, fac : Integer
//algoritma
//	fac <- 1
//	for i <- p downto 1 do
//		fac <- fac * i
//	endfor
//	return fac
//endfunction

//function facY(q : Integer) -> Integer
//kamus
//	i, fac : Integer
//algoritma
//	fac <- 1
//	for i <- q downto 1 do
//		fac <- fac * i
//	endfor
//	return fac
//endfunction

//function mut(p, q : Integer) -> Integer
//kamus
//	i, fac, mutasi : Integer
//algoritma
//	fac <- 1
//	if p >= q then
//		for i <- p-q downto 1 do
//			fac <- fac * i
//		endfor
//		mutasi <- facX(p) div fac
//	else then
//		for i <- q-p downto 1 do
//			fac <- fac * i
//		endfor
//		mutasi <- facY(q) div fac
//	endif
//	return mutasi
//endfunction

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
