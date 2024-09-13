package main

import (
	"fmt"
	"strconv"
)

func main() {

	var bilBul string
	var hasil int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilBul)

	hasil = 0

	for i := 0; i < len(bilBul); i++ {

		num, _ := strconv.Atoi(bilBul[i : i+1])
		hasil += num

	}

	fmt.Println(hasil)

}
