package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	cetak_rerata(N)
}

func cetak_rerata(N int) {
	rerata(N, 0, 1)
}

func rerata(N, total, i int) {
	if i > N {
		return
	}
	total += i
	rata := float64(total) / float64(i)
	fmt.Println(total, rata)
	rerata(N, total, i+1)
}
