// package main

// import "fmt"

// func main() {
// 	var s string
// 	fmt.Scan(&s)
// 	capitalize(s, len(s))
// }

// func capitalize(s string, n int) {
// 	if n > 0 {
// 		capitalize(s, n-1)
// 		fmt.Print(string(s[n-1] - 32))
// 	}
// }

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	lowercase(s, len(s))
}

func lowercase(s string, n int) {
	if n > 0 {
		lowercase(s, n-1)
		fmt.Print(string(s[n-1] + 32))
	}
}
