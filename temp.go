package main

import "fmt"

func main() {
	var s1, s2, s3, s4, s5 float32
	fmt.Scan(&s1)
	fmt.Scan(&s2)
	fmt.Scan(&s3)
	fmt.Scan(&s4)
	fmt.Scan(&s5)

	if s1 < s2 && s2 < s3 && s3 < s4 && s4 < s5 {
		fmt.Println("Stabil Naik")
	} else if s1 > s2 && s2 > s3 && s3 > s4 && s4 > s5 {
		fmt.Println("Stabil Turun")
	} else {
		fmt.Println("Tidak Stabil")
	}
}
