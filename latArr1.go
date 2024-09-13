package main

import "fmt"

const N int = 100

type tab [N]rune

func main() {
	var T tab
	var n int
	masukanArray(&T, &n)
	fmt.Println(isPalindrom(T, n))
}

func masukanArray(T *tab, n *int) {
	var x rune
	fmt.Scanf("%c", &x)
	for x != '.' && *n <= N {
		T[*n] = x
		*n += 1
		fmt.Scanf("%c", &x)
	}
}

func reverseArray(T *tab, n int) {
	var temp rune
	var i, j int
	i = 0
	j = n - 1
	for i < j {
		temp = T[i]
		T[i] = T[j]
		T[j] = temp
		i += 1
		j -= 1
	}
}

func isPalindrom(T tab, n int) bool {
	var x tab
	var hasil bool
	hasil = true
	for i := 0; i < n; i++ {
		x[i] = T[i]
	}

	reverseArray(&x, n)
	for i := 0; i < n; i++ {
		if x[i] != T[i] {
			hasil = false
		}
	}

	return hasil
}
