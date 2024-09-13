package main

import "fmt"

func main() {
	var n, gol, loss, win, draw, lose, totGol, totLoss, selisih, point int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&gol, &loss)
		hitungMenang(gol, loss, &win)
		hitungDraw(gol, loss, &draw)
		hitungKalah(gol, loss, &lose)
		hitungJumGolKegolanSelisih(gol, loss, &totGol, &totLoss, &selisih)
	}
	hitungJumPoint(win, draw, &point)
	fmt.Println(n, win, draw, lose, totGol, totLoss, selisih, point)
}

func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm++
	}
}

func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd++
	}
}

func hitungKalah(g, k int, jk *int) {
	if g < k {
		*jk++
	}
}

func hitungJumGolKegolanSelisih(g, k int, jg, jk, jsg *int) {
	*jg += g
	*jk += k
	*jsg = *jg - *jk
}

func hitungJumPoint(jm, jd int, jp *int) {
	*jp = (jm * 3) + (jd)
}
