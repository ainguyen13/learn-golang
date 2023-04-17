package main

import "fmt"

func checkTriangle (a float64, b float64, c float64) bool {
	if (a + b > c && b + c > a && c + a > b){
		return true
	}
	return false
}

func main(){
	fmt.Print(checkTriangle(3,4,15))
}