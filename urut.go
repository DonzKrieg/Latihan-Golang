package main

import "fmt"

func urutkan(x, y *int) {
	if *x > *y {
		*x, *y = *y, *x
	}
}

func main() {
	var x, y int
	fmt.Scanln(&x, &y)
	urutkan(&x, &y)
	fmt.Println(x, y)

}

// procedure urutkan(in/out x,y:Integer)
// algoritma
// 	if x > y then
// 		x, y <- y, x
// 	endif
// endprocedure

// Program Terurut
// kamus
// 	x, y: Integer
// algoritma
// 	input(x,y)
// 	urutkan(x,y)
// 	output(x,y)
// endprogram
