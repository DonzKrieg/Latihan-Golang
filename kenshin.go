package main

import "fmt"

func main() {

	var n, pow, spd, ePow, eSpd, i, kill int
	fmt.Scan(&n)
	fmt.Scanln(&pow, &spd)
	kill = 0

	for i = 0; i < n; i++ {
		fmt.Scanln(&ePow, &eSpd)
		if pow < 3 || spd < 3 {
			break
		} else {
			if ePow >= pow && eSpd >= spd {
				kill += 0
			} else if ePow >= pow && eSpd < spd {
				kill += 1
				spd -= 6
			} else if ePow < pow && eSpd >= spd {
				kill += 1
				pow -= 6
			} else if ePow < pow && eSpd < spd && spd > pow {
				kill += 1
				spd -= 6
			} else if ePow < pow && eSpd < spd && pow > spd {
				kill += 1
				pow -= 6
			}
		}
	}
	fmt.Println(kill, pow, spd)
}
