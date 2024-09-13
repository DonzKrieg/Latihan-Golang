package main

import "fmt"

func main() {

	var months = [12]string{

		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"Desember",
	}

	var cut1 = months[2:6]
	fmt.Println(cut1)
	fmt.Println(len(cut1))
	fmt.Println(cap(cut1))

	// months[5] = "Monday"
	// fmt.Println(months)

	var cut2 = months[10:]
	fmt.Println(cut2)

	var cut3 = append(cut2, "Break")
	fmt.Println(cut3)

	fmt.Println(cut2)
	fmt.Println(months)

	// lorem ipsum dolor

	newSlice := make([]string, 5, 6)
	newSlice[0] = "Ihsan"
	newSlice[1] = "Romdhon"
	newSlice[2] = "Benaya"
	newSlice[3] = "Abidzar"
	newSlice[4] = "Adistya"
	fmt.Println(newSlice)

}
