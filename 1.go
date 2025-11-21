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
			fmt.Println("Loi,xin hay nhap so.")
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
			fmt.Println("Loi xin hay nhap so.")
		} else {
			break
		}
	}

	if a == 0 && b != 0 {
		fmt.Println("Phuong trinh vo nghiem")
	} else if a == 0 && b == 0 {
		fmt.Println("Phuong trinh co vo so nghiem")
	} else {
		x := -b / a
		fmt.Printf("Vay voi a = %.1f va b = %.1f ta co phuong trinh %.1fx + %.1f va x co ket qua la %.1f\n", a, b, a, b, x)
	}
}
