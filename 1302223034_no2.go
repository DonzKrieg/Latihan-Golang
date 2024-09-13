package main

import (
	"fmt"
	"unicode"
)

func main() {
	var karakter1, karakter2 byte
	fmt.Scanf("%c %c", &karakter1, &karakter2)
	fmt.Println(cekKesamaanKarakter(karakter1, karakter2))
}

func cekKesamaanKarakter(x, y byte) bool {
	if unicode.IsLetter(rune(x)) && unicode.IsLetter(rune(y)) {
		return unicode.ToLower(rune(x)) == unicode.ToLower(rune(y))
	}
	return x == y
}
