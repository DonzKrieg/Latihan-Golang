// Soal No. 6B

// PseudoCode
// Program My_Tel-U2
// kamus
// 		poin : integer
// 		valid : boolean
// algoritma
// 		valid <- true
//		output("Masukkan Poin Anda: ")
// 		input(poin)
// 		if poin < 50 then
// 			valid <- not valid
// 			repeat
// 				output("=== Masukan Tidak Valid ===")
//				output("Masukkan Poin Anda: ")
//				input(poin)
// 			until poin >= 50
// 		endif
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
	var valid bool
	valid = true

	fmt.Print("Masukkan Poin Anda: ")
	fmt.Scan(&poin)

	if poin < 50 {

		valid = !valid
		for !valid {

			fmt.Println("=== Masukan Tidak Valid ===")
			fmt.Print("Masukkan Poin Anda: ")
			fmt.Scan(&poin)
			valid = poin >= 50

		}

	}

	if poin > 200 {

		fmt.Println("You are Gold User")

	} else if poin >= 100 && poin <= 200 {

		fmt.Println("You are Platinum User")

	} else if poin >= 50 && poin <= 99 {

		fmt.Println("You are Silver User")

	}

}
