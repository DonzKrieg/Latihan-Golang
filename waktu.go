package main

import "fmt"

func main() {

	var hour, min, sec, conv int
	fmt.Scanln(&hour, &min, &sec)

	conv = jamConv(hour) + menitConv(min) + sec
	fmt.Println("Hasil konversi:", conv, "Detik")

}

func jamConv(jam int) int {
	return jam * 3600
}

func menitConv(menit int) int {
	return menit * 60
}
