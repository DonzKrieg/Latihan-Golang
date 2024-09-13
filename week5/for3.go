package main

import "fmt"

func main() {

	var n, i, jam, total, avg float32
	fmt.Scan(&n)
	total = 0

	for i = 1; i <= n; i++ {

		fmt.Print("Hari ke-", i, ": ")
		fmt.Scan(&jam)

		total += jam

	}

	avg = total / n
	fmt.Println("Rata-rata:", avg)

}
