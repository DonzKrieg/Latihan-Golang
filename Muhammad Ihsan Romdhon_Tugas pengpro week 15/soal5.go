// Soal No. 5

// PseudoCode
// Program Digit_Terurut
// kamus
// 		digit : string
// 		ascen, descen : boolean
// 		prev, next, i : integer
// algoritma
// 		output("Masukkan Digit: ")
//		input(digit)
// 		for i <- 0 to lenght(digit) - 1 do
// 			prev, _ <- convert string to integer (digit[i : i+1])
// 			next, _ <- convert string to integer (digit[i+1 : i+2])
// 			if ascen and next < prev then
// 				ascen <- false
// 				break
// 			else if descen and next > prev then
// 				descen <- false
// 				break
// 			endif
// 			ascen <- next > prev
//			descen <- next < prev
// 		endfor
// 		if descen and not ascen then
//			output("descending")
//		else if ascen and not descen then
//			output("ascending")
//		else then
//			output("tidak terurut")
// 		endif
// endprogram

// Golang
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var digit string
	var ascen, descen bool
	var prev, next, i int
	fmt.Print("Masukkan Digit: ")
	fmt.Scan(&digit)

	for i = 0; i < len(digit)-1; i++ {
		prev, _ = strconv.Atoi(digit[i : i+1])
		next, _ = strconv.Atoi(digit[i+1 : i+2])

		if ascen && next < prev {
			ascen = false
			break
		} else if descen && next > prev {
			descen = false
			break
		}

		ascen = next > prev
		descen = next < prev
	}

	if descen && !ascen {
		fmt.Println("descending")
	} else if ascen && !descen {
		fmt.Println("ascending")
	} else {
		fmt.Println("tidak terurut")
	}

}
