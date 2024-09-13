package main

import "fmt"

type series struct {
	nama                string
	eps, durasi, totDur int
}

func main() {
	var film series
	var n int
	fmt.Scan(&film.nama, &film.eps, &film.durasi)
	fmt.Scan(&n)

	film.totDur = film.eps * film.durasi

	var hari int
	if film.totDur%(n*60) != 0 {
		hari = film.totDur/(n*60) + 1
	} else {
		hari = film.totDur / (n * 60)
	}

	fmt.Println(hari, "hari")
}
