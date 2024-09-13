package main

import "fmt"

func main() {
	var pilih int
	for {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			hitungJumlah()
			break
		case 2:
			hitungKali()
			break
		case 3:
			hitungBagi()
			break
		}
		if pilih == 4 {
			break
		}
	}
}

func menu() {
	fmt.Println("-----------------------")
	fmt.Println("        M E N U        ")
	fmt.Println("-----------------------")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
	fmt.Println("-----------------------")

}

func hitungJumlah() {
	var a, b int
	fmt.Print("Masukkan dua bilangan yang akan dijumlahkan: ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil penjumlahan:", a+b)
}

func hitungKali() {
	var a, b int
	fmt.Print("Masukkan dua bilangan yang akan dikalikan: ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil perkalian:", a*b)
}

func hitungBagi() {
	var a, b float64
	fmt.Print("Masukkan dua bilangan yang akan dibagikan: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Hasil pembagian: %.1f\n", a/b)
}
