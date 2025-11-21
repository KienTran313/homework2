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
		fmt.Print("Nhap so a: ")
		scanner.Scan()
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		a = number
		if err != nil {
			fmt.Println("Loi,hay nhap so.")
		} else {
			break
		}
	}

	var b float64
	for {
		fmt.Print("Nhap so b: ")
		scanner.Scan()
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		b = number
		if err != nil {
			fmt.Println("Loi,hay nhap so")
		} else {
			break
		}
	}

	for {
		fmt.Print("Nhap phep toan (+,-,*,/) : ")
		scanner.Scan()
		text := scanner.Text()

		x1 := a + b
		if text == "+" {
			fmt.Printf("%.1f + %.1f = %.1f\n", a, b, x1)
			break
		}
		x2 := a - b
		if text == "-" {
			fmt.Printf("%.1f - %.1f = %.1f\n", a, b, x2)
			break
		}
		x3 := a * b
		if text == "*" {
			fmt.Printf("%.1f * %.1f = %.1f\n", a, b, x3)
			break
		}
		x4 := a / b
		if text == "/" {
			fmt.Printf("%.1f / %.1f = %.1f\n", a, b, x4)
			break
		}
	}
}
