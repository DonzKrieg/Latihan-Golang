package main

import "fmt"

func main() {
	var totBagus, totCacat int
	totBagus, totCacat = 0, 0
	countBagusCacat(&totBagus, &totCacat)
	fmt.Println(totBagus, totCacat)
}

func countBagusCacat(totBagus, totCacat *int) {
	for {
		var quality string
		var jumlah int
		fmt.Scan(&quality, &jumlah)
		if quality == "bagus" {
			*totBagus += jumlah
		} else if quality == "cacat" {
			*totCacat += jumlah
		} else {
			break
		}
	}
}
