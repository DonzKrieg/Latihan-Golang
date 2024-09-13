// program Fibonacci
// kamus
// 	bil, hasil : Integer
// algoritma
// 	input(bil)
// 	hasil <- fibo(bil)
// 	output(hasil)
// endprogram

// function fibo(n : Integer) -> Integer
// kamus
// 	i, t1, t2, next, tot : Integer
// algoritma
// 	t1 <- 0
// 	t2 <- 1
// 	tot <- 0
// 	for i = 1 <- i to n, i++ do
// 		if i == 1 then
// 			tot <- tot + t1
// 			continue
//		endif
// 		if i == 2 then
// 			tot <- tot + t2
// 			continue
//		endif
// 		next <- t1 + t2
// 		t1 <- t2
// 		t2 <- next
// 		tot <- tot + next
//  endfor
// 	return tot
// endfuncttion

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
