package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	rekursif(N)
}

func rekursif(N int) {
	var hasil, i int
	if N == 1 {
		hasil = 1
	} else {
		for i = 1; i <= N; i++ {
			hasil += i
		}
		fmt.Println(hasil)
	}
}
