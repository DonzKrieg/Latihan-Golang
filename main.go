package main

import (
	"fmt"
	"strconv"
)

func soal1() {
	var username, password string
	var percobaan int = 0

	fmt.Print("Masukkan username: ")
	fmt.Scan(&username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&password)

	for username != "admin" || password != "admin" {
		fmt.Println("==== USERNAME ATAU PASSWORD SALAH ====")
		fmt.Print("Masukkan username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&password)
		percobaan++
	}
	fmt.Printf("%d percobaan gagal\n", percobaan)
	fmt.Printf("Login berhasil\n")
}

func soal2() {
	var saldo, pengeluaran int

	fmt.Print("Masukkan Saldo: ")
	fmt.Scan(&saldo)

	fmt.Printf("Masukkan pengeluaran: ")
	fmt.Scan(&pengeluaran)

	for pengeluaran != 0 {
		saldo = saldo - pengeluaran

		fmt.Printf("Masukkan pengeluaran: ")
		fmt.Scan(&pengeluaran)
	}
	fmt.Printf("Sisa saldo: %d", saldo)
}

func soal3() {
	var input string
	var total int = 0

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&input)

	// 1 2 3 4 5
	for i := len(input) - 1; i >= 0; i-- {
		x, _ := strconv.Atoi(input[i : i+1])
		total += x
		fmt.Printf("%d ", x)
	}
	fmt.Print("\n", total)
}

func soal4() {
	var totalCangkir int = 0

	var persediaanGula, persediaanKopi int
	var butuhGula, butuhKopi int

	fmt.Scanln(&persediaanGula, &persediaanKopi, &butuhGula, &butuhKopi)

	for butuhGula <= persediaanGula && butuhKopi <= persediaanKopi {
		persediaanGula -= butuhGula
		persediaanKopi -= butuhKopi
		totalCangkir++
	}
	fmt.Print(totalCangkir)
}

func soal5() {
	var input string
	var valid bool = true

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&input)

	// 1010101010
	// 0123456789

	for i := 0; i < len(input)-1; i++ {
		if valid == false {
			break
		}
		prev, _ := strconv.Atoi(input[i : i+1])
		next, _ := strconv.Atoi(input[i+1 : i+2])
		valid = next == prev+1 || next == prev-1
	}
	fmt.Print(valid)
}

func soal6() {
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

func main() {
	//soal1()
	soal2()
	//soal3()
	//soal4()
	//soal5()
	//soal6()
}
