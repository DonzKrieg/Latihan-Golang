// Pseudo Code

// program ManagerEPL
// kamus
// 		match1, match2, match3, match4, match5 : String
// algoritma
// 		input (match1, match2, match3, match4, match5)
//		if match1 = "kalah" and match2 = "kalah" and match3 = "kalah" and match4 = "kalah" and match5 = "kalah" then
//			output("Dipecat!!!")
//		else
//			output("Tidak Dipecat")
//		endif
// endprogram

package main

import "fmt"

func main() {

	var match1, match2, match3, match4, match5 string
	fmt.Scan(&match1, &match2, &match3, &match4, &match5)

	if match1 == "kalah" && match2 == "kalah" && match3 == "kalah" && match4 == "kalah" && match5 == "kalah" {
		fmt.Println("Dipecat!!!")
	} else {
		fmt.Println("Tidak Dipecat")
	}

}
