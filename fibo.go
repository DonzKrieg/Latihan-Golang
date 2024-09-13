package main

import "fmt"

func main() {

	var bil, hasil int
	fmt.Scan(&bil)
	hasil = fibo(bil)
	fmt.Println(hasil)

}

func fibo(n int) int {
	var i, t1, t2, next, tot int
	t1 = 0
	t2 = 1
	tot = 0
	for i = 1; i <= n; i++ {
		if i == 1 {
			tot += t1
			continue
		}
		if i == 2 {
			tot += t2
			continue
		}
		next = t1 + t2
		t1 = t2
		t2 = next
		tot += next

	}
	return tot
}
