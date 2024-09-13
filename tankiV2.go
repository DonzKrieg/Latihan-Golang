package main

import "fmt"

func main() {
	var tanki, ember, jumlah int
	isiTanki(tanki, ember, jumlah)
}

func isiTanki(tanki, ember, jumlah int) {
	fmt.Scan(&tanki)
	jumlah = 0
	for {
		fmt.Scan(&ember)
		jumlah += ember
		fmt.Println(jumlah >= tanki)
		if jumlah >= tanki {
			break
		}
	}
}
