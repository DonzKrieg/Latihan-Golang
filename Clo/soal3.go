package main

import "fmt"

func main() {

	var bil1, bil2, bil3, bil4, bil5 int
	fmt.Println("Masukkan Bilangan pertama dan Kedua: ")
	fmt.Scan(&bil1, &bil2)

	bil3 = bil2 + bil1
	bil4 = bil3 + bil2
	bil5 = bil4 + bil3

	fmt.Println("Hasil", bil3, bil4, bil5)

}
