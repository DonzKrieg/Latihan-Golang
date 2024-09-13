// Soal No. 6A

// PseudoCode
// Program My_Tel-U1
// kamus
// 		poin : integer
// algoritma
//		output("Masukkan Poin Anda: ")
// 		input(poin)
// 		if poin > 200 then
// 			output("You are Gold User")
// 		else if poin >= 100 and poin <= 200 then
// 			output("You are Platinum User")
// 		else if poin >= 50 and poin <= 99 then
// 			output("You are Silver User")
// 		endif
// endprogram

// Golang
package main

import "fmt"

func main() {

	var poin int
	fmt.Print("Masukkan Poin Anda: ")
	fmt.Scan(&poin)

	if poin > 200 {

		fmt.Println("You are Gold User")

	} else if poin >= 100 && poin <= 200 {

		fmt.Println("You are Platinum User")

	} else if poin >= 50 && poin <= 99 {

		fmt.Println("You are Silver User")

	}

}
