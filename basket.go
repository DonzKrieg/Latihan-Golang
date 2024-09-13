package main

import "fmt"

func main() {

	var blue, red, i, totPoin, totGol, totLoss, selGol int

	totPoin = 0
	totGol = 0
	totLoss = 0
	selGol = 0

	for i = 1; i <= 38; i++ {
		fmt.Scanln(&blue, &red)
		if blue > red {
			totGol += blue
			totLoss += red
			totPoin += 3
		} else if blue < red {
			totLoss += red
			totGol += blue
		} else {
			totPoin += 1
			totGol += blue
			totLoss += red
		}
	}

	selGol = totGol - totLoss
	fmt.Println(totPoin, totGol, totLoss, selGol)

}
