package main

import "fmt"

func main() {

	var tabPertama, n, jumlah int
	fmt.Scan(&tabPertama, &n)
	jumlah = uangKakak(tabPertama, n)
	fmt.Println(jumlah)

}

func uangKakak(x, y int) int {
	var total, i int
	for i = 1; i <= y; i++ {
		total += x
		x += 2500
	}
	return total
}
