package main

import "fmt"

func main() {

	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(pertemuan(x, y))

}

func pertemuan(a, b int) int {
	var i, tot int
	if a > 0 && b > 0 && a < b {
		for i = 1; i <= 365; i++ {
			if i%a == 0 || i%b == 0 {
				tot += 1
			}
		}
	}
	return tot
}
