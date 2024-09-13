package main

import "fmt"

func findMidPoint(bil []int) bool {
	// bubble sort
	for i := 0; i < len(bil)-1; i++ {
		for j := 0; j < len(bil)-1-i; j++ {
			if bil[j] > bil[j+1] {
				temp := bil[j]
				bil[j] = bil[j+1]
				bil[j+1] = temp
			}
		}
	}
	smallNumber := bil[0]
	midNumber := bil[1]
	largeNumber := bil[2]

	var num []int
	for i := smallNumber; i <= largeNumber; i++ {
		num = append(num, i)
	}

	midIndex := int(len(num) / 2)
	if len(num)%2 == 0 {
		x := num[(num[midIndex-1]+num[midIndex])/2]
		return midNumber == x
	} else {
		x := num[midIndex]
		return midNumber == x
	}
}

func main() {
	var bil []int
	var a, b, c int
	fmt.Printf("Masukkan 3 bilangan: \n")
	fmt.Scan(&a, &b, &c)
	bil = append(bil, a)
	bil = append(bil, b)
	bil = append(bil, c)

	fmt.Println(findMidPoint(bil))
}
