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
	var sign string
	for {
		fmt.Print("Nhap phep toan (+,-,*,/) : ")
		scanner.Scan()
		text := scanner.Text()
		if text == "+" || text == "-" || text == "*" || text == "/" {
			sign = text
			break
		} else {
			fmt.Println("Loi,chi nhap cac dau (+,-,*,/). ")
		}
	}

	var result float64
	for {
		fmt.Printf("Nhap ket qua cua %v %v %v : ", a, sign, b)
		scanner.Scan()
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		result = number
		if err != nil {
			fmt.Println("Loi,hay nhap so.")
		} else {
			break
		}
	}

	var corect float64
	if sign == "+" {
		corect = a + b
	} else if sign == "-" {
		corect = a - b
	} else if sign == "*" {
		corect = a * b
	} else if sign == "/" {
		corect = a / b
		if b == 0 {
			fmt.Println("Khong co ket qua.")
			return
		}
	}

	if result == corect {
		fmt.Println("Ban da lam dung.")
	} else {
		fmt.Printf("Ban da tinh sai,ket qua cua %v %v %v = %v\n", a, sign, b, corect)
	}
}
