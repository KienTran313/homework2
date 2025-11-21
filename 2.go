package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var a float64
	for {
		fmt.Print("Nhap do dai canh a: ")
		scanner.Scan()
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		a = number
		if err != nil {
			fmt.Println("Loi,hay nhap so.")
		} else if number <= 0 {
			fmt.Println("Loi,hay nhap so duong lon hon 0.")
		} else {
			break
		}
	}

	var b float64
	for {
		fmt.Print("Nhap do dai canh b: ")
		scanner.Scan()
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		b = number
		if err != nil {
			fmt.Println("Loi,hay nhap so.")
		} else if number <= 0 {
			fmt.Println("Loi,hay nhap so duong lon hon 0.")
		} else {
			break
		}
	}

	var c float64
	for {
		fmt.Print("Nhap do dai canh c: ")
		scanner.Scan()
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		c = number
		if err != nil {
			fmt.Println("Loi,hay nhap so.")
		} else if number < 0 {
			fmt.Println("Loi,hay nhap so duong.")
		} else {
			break
		}
	}

	if a+b <= c || a+c <= b || b+c <= a {
		// khong hieu vi sao khi them dau "=" vao ket qua lai khac.
		fmt.Println("Day khong phai la tam giac.")
	} else if a*a == c*c+b*b || b*b == a*a+c*c || c*c == a*a+b*b {
		fmt.Println("Day la tam giac vuong.")
	} else {
		fmt.Println("Day la tam giac thuong.")
	}
}
