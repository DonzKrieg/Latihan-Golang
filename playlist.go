package main

import (
	"fmt"
)

type lagu struct {
	judul  string
	artis  string
	durasi int
}

type playlist struct {
	laguList    [20]lagu
	totalDurasi int
	jumlahLagu  int
}

func insertLagu(pl *playlist, n int) {
	// I.S terdefinisi playlist kosong dan n lagu yang ingin dimasukkan.
	// F.S playlist terisi dengan n lagu
	for i := 0; i < n; i++ {
		var l lagu
		fmt.Scanln(&l.judul, &l.artis, &l.durasi)
		if !isMember(*pl, l) {
			pl.laguList[pl.jumlahLagu] = l
			pl.totalDurasi += l.durasi
			pl.jumlahLagu++
		}
	}
}

func isMember(pl playlist, la lagu) bool {
	// I.S Terdefinisi playlist dan sebuah lagu
	// F.S mengembalikan true jika lagu sudah termasuk dalam playlist
	for i := 0; i < pl.jumlahLagu; i++ {
		if pl.laguList[i].judul == la.judul && pl.laguList[i].artis == la.artis {
			return true
		}
	}
	return false
}

func gabungPlaylist(pl1 playlist, pl2 playlist, pl3 *playlist) {
	// I.S Terdifinisi playlist zaki dan playlist mifta yang ingin digabung
	// F.S Playlist zaki dan playlist mifta tergabung pada playlist 3, Note : Playlist Zaki di bagian atas array.
	for i := 0; i < pl1.jumlahLagu; i++ {
		pl3.laguList[i] = pl1.laguList[i]
	}
	pl3.jumlahLagu = pl1.jumlahLagu
	pl3.totalDurasi = pl1.totalDurasi

	for i := 0; i < pl2.jumlahLagu; i++ {
		if !isMember(*pl3, pl2.laguList[i]) {
			pl3.laguList[pl3.jumlahLagu] = pl2.laguList[i]
			pl3.jumlahLagu++
			pl3.totalDurasi += pl2.laguList[i].durasi
		}
	}
}

func tampilPlaylist(pl playlist) {
	// I.S Terdefinisi playlist
	// F.S Menampilkan judul dan artis semua lagu dan total durasi playlist
	for i := 0; i < pl.jumlahLagu; i++ {
		fmt.Printf("%s %s\n", pl.laguList[i].judul, pl.laguList[i].artis)
	}
	durasiJam := pl.totalDurasi / 3600
	durasiMenit := (pl.totalDurasi % 3600) / 60
	durasiDetik := pl.totalDurasi % 60
	fmt.Printf("Durasi Playlist: %02d:%02d:%02d\n", durasiJam, durasiMenit, durasiDetik)
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
