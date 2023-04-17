package main

import "fmt"

func main() {
	fmt.Println(fibonacci2(10))
}
func fibonacci2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	} else {
		fib := make([]int, n+1)
		fib[0] = 0
		fib[1] = 1
		for i := 2; i < n+1; i++ {
			fib[i] = fib[i-1] + fib[i-2]
		}
		return fib[n]
	}
}

// func fibonacci2(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	return fibonacci2(n - 1) + fibonacci2(n - 2)
// }