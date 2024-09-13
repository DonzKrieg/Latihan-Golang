package main

import "fmt"

type chess struct {
	pion, from, to string
	nilai          int
}

func main() {
	var l1, l2, l3 chess
	fmt.Scan(&l1.pion, &l1.from, &l1.to, &l1.nilai)
	fmt.Scan(&l2.pion, &l2.from, &l2.to, &l2.nilai)
	fmt.Scan(&l3.pion, &l3.from, &l3.to, &l3.nilai)

	if l1.nilai > l2.nilai && l1.nilai > l3.nilai {
		fmt.Println(l1.pion, l1.from, "-", l1.to)
	} else if l2.nilai > l1.nilai && l2.nilai > l3.nilai {
		fmt.Println(l2.pion, l2.from, "-", l2.to)
	} else {
		fmt.Println(l3.pion, l3.from, "-", l3.to)
	}
}
