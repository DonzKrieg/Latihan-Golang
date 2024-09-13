package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	fmt.Println(isPrime(bilangan))
}

func isPrime(x int) bool {
	var i, count int
	for i = 2; i < x; i++ {
		if x%i == 0 {
			count += 1
		}
	}
	if count == 1 {
		return true
	} else {
		return false
	}
}
