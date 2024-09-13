package main

import "fmt"

type ChessMove struct {
	Piece string
	From  string
	To    string
	Value int
}

func main() {
	// Input data langkah
	langkah := make([]ChessMove, 3)
	for i := 0; i < 3; i++ {
		fmt.Printf("Masukkan langkah %d (nama buah, petak asal, petak tujuan, nilai): ", i+1)
		fmt.Scan(&langkah[i].Piece, &langkah[i].From, &langkah[i].To, &langkah[i].Value)
	}

	// Mencari langkah dengan nilai tertinggi
	langkahTertinggi := cariLangkahTertinggi(langkah)

	// Menampilkan langkah dengan nilai tertinggi
	fmt.Printf("Langkah dengan nilai tertinggi: %s %s-%s\n", langkahTertinggi.Piece, langkahTertinggi.From, langkahTertinggi.To)
}

func cariLangkahTertinggi(langkah []ChessMove) ChessMove {
	maxNilai := langkah[0].Value
	langkahTertinggi := langkah[0]

	// Mencari langkah dengan nilai tertinggi
	for i := 1; i < len(langkah); i++ {
		if langkah[i].Value > maxNilai {
			maxNilai = langkah[i].Value
			langkahTertinggi = langkah[i]
		}
	}

	return langkahTertinggi
}