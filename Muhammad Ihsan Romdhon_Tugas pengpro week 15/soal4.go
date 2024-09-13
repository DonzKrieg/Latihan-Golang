// Soal No. 4

// PseudoCode
// Program Mobil_Balap
// kamus
// 		finish, win : string
// 		a, b, i : integer
// algoritma
// 		a <- 0
// 		b <- 0
// 		for i <- 1 to 10 do
// 			output("Finish ke-", i, ": ")
//			input(finish)
// 			if finish = "A" then
// 				a <- a + 1
// 			else if finish = "B" then
// 				b <- b + 1
// 			endif
// 			if a = 4 then
// 				win <- "Pemenangnya adalah B !!!"
// 			else if b = 4 then
// 				win <- "Pemenangnya adalah A !!!"
// 			endif
// 		endfor
// 		if a <= 5 and b <= 5 then
// 			output(win)
// 		else then
// 			output("===Invalid===")
// 		endif
// endprogram

// Golang
package main

import "fmt"

func main() {

	var finish, win string
	var a, b, i int

	a = 0
	b = 0

	for i = 1; i <= 10; i++ {

		fmt.Print("Finish ke-", i, ": ")
		fmt.Scan(&finish)

		if finish == "A" {

			a = a + 1

		} else if finish == "B" {

			b = b + 1

		}

		if a == 4 {

			win = "Pemenangnya adalah B !!!"

		} else if b == 4 {

			win = "Pemenangnya adalah A !!!"

		}

	}

	if a <= 5 && b <= 5 {

		fmt.Println(win)

	} else {

		fmt.Println("===Invalid===")

	}

}
