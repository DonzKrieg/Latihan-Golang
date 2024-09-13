// program Belanja
// kamus
//     belanja, diskon : real
//     member : bool
// algoritma
//     input(belanja, member)
//     diskon <- total(belanja, member)
//     output(diskon)
// endprogram

// function total(x : real, y : boolean) -> real
// kamus
// 		hasil : real
// algoritma
// 		if x > 100000 and y == true then
//			hasil <- x - (x * 0.1)
//		else if x > 100000 and y == false then
//			hasil <- x - (x * 0.05)
//		else then
//			hasil <- x
//		endif
//		return hasil
// endfunction

package main

import "fmt"

func main() {

	var belanja, diskon float32
	var member bool
	fmt.Scanln(&belanja, &member)

	diskon = total(belanja, member)
	fmt.Println(diskon)

}

func total(x float32, y bool) float32 {
	var hasil float32
	if x > 100000 && y == true {
		hasil = x - (x * 0.1)
	} else if x > 100000 && y == false {
		hasil = x - (x * 0.05)
	} else {
		hasil = x
	}
	return hasil
}
