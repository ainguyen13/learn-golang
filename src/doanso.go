package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(101)

	var guess int 
	for {
		fmt.Print("Moi ban nhap so: ")
		fmt.Scan(&guess)

		if guess > x {
			fmt.Println("Lon hon x")
		}
		if guess < x {
			fmt.Println("Nho hon x")
		} else
		{
			fmt.Println("Bang x")
			break
		}
		
	}




















	