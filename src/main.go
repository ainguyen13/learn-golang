package main

import (
	"fmt"
)
func main() {
    fmt.Println(factorial(5))
	fmt.Println(fibonacci(5))
}

/*
TÌm giai thừa của một số
*/
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n - 1)
}
/*
Hàm tính chuỗi fibonacci của một số
*/
func fibonacci(n int) int {
	if n == 1 {
		return 1
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}