// program WaktuConverter
// kamus
//     hour, min, sec, conv : Integer
// algoritma
// 	input(hour, min, sec)
// 	conv <- jamConv(hour) + menitConv(min) + sec
// 	output("Hasil konversi:", conv, "Detik")
// endprogram

// function jamConv(jam: Integer) -> Integer
// algoritma
//  return jam*3600
// endfunction

// function menitConv(menit: Integer) -> Integer
// algoritma
//  return menit*60
// endfunction

// package main

// import "fmt"

// func main() {

// 	var hour, min, sec, conv int
// 	fmt.Scanln(&hour, &min, &sec)

// 	conv = jamConv(hour) + menitConv(min) + sec
// 	fmt.Println("Hasil konversi:", conv, "Detik")

// }

// func jamConv(jam int) int {
// 	return jam * 3600
// }

// func menitConv(menit int) int {
// 	return menit * 60
// }

package main

import "fmt"

func main() {
	var jam, menit, detik int
	fmt.Scan(&jam, &menit, &detik)
	fmt.Println(convJam(jam) + convMen(menit) + detik)
}

func convJam(jam int) int {
	return jam * 3600
}

func convMen(jam int) int {
	return jam * 60
}
