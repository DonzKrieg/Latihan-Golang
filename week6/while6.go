package main

import "fmt"

func main() {

	var kapasitasTanki, isiTanki, literEmber int
	var tankiPenuh bool = false

	fmt.Print("Masukkan kapasitas tanki (liter): ")
	fmt.Scan(&kapasitasTanki)

	fmt.Print("Masukkan air (liter): ")
	fmt.Scan(&literEmber)

	for literEmber <= kapasitasTanki {
		isiTanki += literEmber
		tankiPenuh = isiTanki >= kapasitasTanki
		fmt.Println(tankiPenuh)

		if tankiPenuh {
			break
		}

		fmt.Print("Masukkan air (liter): ")
		fmt.Scan(&literEmber)
	}

}
