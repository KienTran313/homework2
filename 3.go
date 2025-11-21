package main

import (
	"bufio"
	"fmt"
	"math"
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
			fmt.Println("Loi,xin hay nhap so.")
		} else {
			break
		}
	}

	var c float64
	for {
		fmt.Print("Nhap so c: ")
		scanner.Scan()
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		c = number
		if err != nil {
			fmt.Println("Loi,xin hay nhap so.")
		} else {
			break
		}
	}
	fmt.Printf("Vay ta co phuong trinh %.1fx2 + %.1fx + %.1f = 0\n", a, b, c)
	if a == 0 {
		fmt.Println("Phuong trinh vo nghiem.")
		return
	}
	delta := b*b - 4*a*c
	if delta < 0 {
		fmt.Println("Phuong trinh vo nghiem.")
	} else if delta == 0 {
		x := -b / 2 * a
		fmt.Printf("Phuong trinh co nghiem la x1 = x2 = %.1f\n ", x)
	} else {
		x1 := -b + math.Sqrt(delta)/2*a
		x2 := -b - math.Sqrt(delta)/2*a
		fmt.Printf("Phuong trinh co hai nghiem la x1 = %.1f va x2 = %.1f\n", x1, x2)
	}
}
