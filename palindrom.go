package main

import "fmt"

const N int = 100

type tab [N]rune

func masukanArray(T *tab, n *int) {
	var char rune
	for {
		fmt.Scanf("%c", &char)
		if char == '.' {
			break
		}
		T[*n] = char
		*n++
	}
}

func reverseArray(T *tab, n int) {
	for i := 0; i < n/2; i++ {
		T[i], T[n-i-1] = T[n-i-1], T[i]
	}
}

func isPalindrom(T tab, n int) bool {
	reverseArray(&T, n)
	for i := 0; i < n; i++ {
		if T[i] != T[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var T tab
	var n int
	masukanArray(&T, &n)
	fmt.Println(isPalindrom(T, n))
}
