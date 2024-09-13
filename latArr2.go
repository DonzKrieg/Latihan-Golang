package main
import "fmt"


type lagu struct {
	judul, artis string
	durasi int
}

type playlist struct {
	lagulist [20]lagu
	totDurasi, jumLagu int
}

func main() {
	var pl1, pl2, pl3 playlist
	var nZaki, nMifta int
	fmt.Scanln(&nZaki)
	insertLagu(&pl1, nZaki)
	fmt.Scanln(&nMifta)
	insertLagu(&pl2, nMifta)
	gabungPlaylist(pl1, pl2, &pl3)
	tampilPlaylist(pl3)
}

func insertLagu(pl *playlist, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&pl.lagulist[i].judul, &pl.lagulist[i].artis, &pl.lagulist[i].durasi)
		pl.jumLagu += 1
		pl.totDurasi += pl.lagulist[i].durasi
	}
}

func isMember(pl playlist, la lagu) bool {
	for i := 0; i < pl.jumLagu {
		if pl.lagulist[i].judul
	}
}