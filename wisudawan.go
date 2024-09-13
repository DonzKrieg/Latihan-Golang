package main

import (
	"fmt"
)

const nmax = 1000

type Wisudawan struct {
	Nama     string
	NIM      string
	Eprt     int
	Semester int
	Ipk      float64
}

func isiDataWisudawan(wisudawan *[]Wisudawan) {
	for i := 0; i < nmax; i++ {
		var data Wisudawan
		fmt.Print("Masukkan nama: ")
		fmt.Scanln(&data.Nama)
		fmt.Print("Masukkan NIM (none untuk berhenti): ")
		fmt.Scanln(&data.NIM)
		if data.NIM == "none" {
			if i == 0 {
				fmt.Println("Tidak ada data wisudawan.")
			}
			break
		}
		fmt.Print("Masukkan eprt: ")
		fmt.Scanln(&data.Eprt)
		fmt.Print("Masukkan jumlah semester: ")
		fmt.Scanln(&data.Semester)
		fmt.Print("Masukkan IPK: ")
		fmt.Scanln(&data.Ipk)

		*wisudawan = append(*wisudawan, data)
	}
}

func eprtTertinggi(wisudawan []Wisudawan) int {
	maxEprt := wisudawan[0].Eprt
	for _, w := range wisudawan {
		if w.Eprt > maxEprt {
			maxEprt = w.Eprt
		}
	}
	return maxEprt
}

func ipkTerendah(wisudawan []Wisudawan) float64 {
	minIpk := wisudawan[0].Ipk
	for _, w := range wisudawan {
		if w.Ipk < minIpk {
			minIpk = w.Ipk
		}
	}
	return minIpk
}

func rataRataSemester(wisudawan []Wisudawan) float64 {
	totalSemester := 0
	for _, w := range wisudawan {
		totalSemester += w.Semester
	}
	return float64(totalSemester) / float64(len(wisudawan))
}

func main() {
	var dataWisudawan []Wisudawan

	isiDataWisudawan(&dataWisudawan)

	if len(dataWisudawan) == 0 {
		return
	}

	fmt.Printf("Eprt tertinggi: %d\n", eprtTertinggi(dataWisudawan))
	fmt.Printf("IPK terendah: %.2f\n", ipkTerendah(dataWisudawan))
	fmt.Printf("Rata-rata semester: %.2f\n", rataRataSemester(dataWisudawan))
}
