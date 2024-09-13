package main

import "fmt"

func main() {
	var str string
	var n int
	fmt.Scan(&str)
	n = len(str)
	reverse(str, n)
}

func reverse(str string, n int) {
	if n > 0 {
		fmt.Print(string(str[n-1]))
		reverse(str, n-1)
	}
}
