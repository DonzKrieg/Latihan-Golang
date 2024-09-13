package main

import "fmt"

func main() {

	var hour, min, sec int
	fmt.Print("Masukkan detik: ")
	fmt.Scan(&sec)

	hour = sec / 3600
	min = (sec / 60) % 60
	sec = sec % 60

	fmt.Printf("%d Jam, %d Menit, %d Detik", hour, min, sec)

}
