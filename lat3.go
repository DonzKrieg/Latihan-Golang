package main

import "fmt"

type hasilGanjil struct {
	bilAkhir, jumTot int
	rata             float64
}

func main() {
	var n int
	fmt.Scan(&n)
	hasil := jumBil(n)
	fmt.Println(hasil.bilAkhir, hasil.jumTot, hasil.rata)
}

func jumBil(n int) hasilGanjil {
	var bil, jum int
	for i := 0; i < n; i++ {
		bil = 2*i + 1
		jum += bil
	}

	var hg hasilGanjil
	hg.bilAkhir = bil
	hg.jumTot = jum

	hg.rata = float64(hg.jumTot) / float64(n)
	return hg
}
