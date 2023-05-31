package main

import(
	"fmt"
)
func main() {
	a := removeItemSliceKeepOrder([]string{"a", "b", "c", "d", "e", "f"}, 2)
	fmt.Println(a)
}