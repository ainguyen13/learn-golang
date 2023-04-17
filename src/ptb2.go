package main

import (
    "fmt"
    "math"
)

func main() {
    var a, b, c float64
    fmt.Println("Nhập a: ")
    fmt.Scan(&a)
    fmt.Println("Nhập b: ")
    fmt.Scan(&b)
    fmt.Println("Nhập c: ")
    fmt.Scan(&c)
	if a == 0 {
		x := -c /b
		fmt.Println("Phuong trinh co 1 nghiem duy nhat: ",x )
	}
    delta := b*b - 4*a*c
	if delta == 0 {
		xx := -b/(2*a)
		fmt.Println("Phuong trinh co mot nghiem duy nhat: ", xx)
	}
    if delta > 0 {
        x1 := (-b + math.Sqrt(delta)) / (2 * a)
        x2 := (-b - math.Sqrt(delta)) / (2 * a)
        fmt.Printf("Phương trình có 2 nghiệm thực: x1 = %.2f và x2 = %.2f", x1, x2)
    } else {
        realPart := -b / (2 * a)
        imaginaryPart := math.Sqrt(-delta) / (2 * a)
		fmt.Println("Ở đây delta < 0 thì phải vô nghiệm nhưng do y/c bài toán delta < 0 thì phải trả kết quả về số phức")
        fmt.Printf("Phương trình có 2 nghiệm phức: x1 = %.2f+%.2fi và x2 = %.2f-%.2fi", realPart, imaginaryPart, realPart, imaginaryPart)
    }
}
