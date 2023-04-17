package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		err    error
		height float64
		weight float64
	)
	height, err = readNumberFromKeyboard("Chieu cao cua ban: ")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	weight, err = readNumberFromKeyboard("Can nang cua ban: ")
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	bmi := CalculateBMI(height, weight)
	fmt.Println("Chỉ số bmi của bạn = ", bmi)
	
	switch {
	case bmi < 16:
		fmt.Println("Severe thinness")
	case bmi < 16.9:
		fmt.Println("Moderate thinness")
	case bmi < 18.4:
		fmt.Println("Mild thinness")
	case bmi < 24.9:
		fmt.Println("Normal")
	case bmi < 29.9:
		fmt.Println("Pre-obese")
	case bmi < 34.9:
		fmt.Println("Obese (Class I)")
	case bmi < 39.9:
		fmt.Println("Obese (Class II)")
	default:
		fmt.Println("Obese (Class III)")
	}
}
/*
Đọc từ bàn phím và convert string sang float
*/
func readNumberFromKeyboard(msg string) (result float64, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")
	if result, err = strconv.ParseFloat(str, 64); err != nil {
		return 0.0, err
	}
	return result, nil
}

func CalculateBMI(height float64, weight float64) (index float64) {
	return  weight / (height * height)
}
