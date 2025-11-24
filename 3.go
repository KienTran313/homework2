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
	if a == 0 {
		fmt.Printf("Vay ta co phuong trinh %vx + %v = 0\n", b, c)
		if b == 0 && c != 0 {
			fmt.Println("Phuong trinh vo nghiem. ")
			return
		} else if b == 0 && c == 0 {
			fmt.Println("Phuong trinh co vo so nghiem.")
			return
		} else if b != 0 {
			x3 := -c / b
			fmt.Printf("Phuong trinh co mot nghiem x = %v\n", x3)
			return
		}
	}

	fmt.Printf("Vay ta co phuong trinh %vx^2 + %vx + %v = 0\n", a, b, c)
	delta := (b * b) - (4 * a * c)
	if delta < 0 {
		fmt.Println("Phuong trinh vo nghiem.")
	} else if delta == 0 {
		x := -b / (2 * a)
		fmt.Printf("Phuong trinh co nghiem la x1 = x2 = %v\n ", x)
	} else {
		x1 := (-b + math.Sqrt(delta)) / (2 * a)
		x2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Printf("Phuong trinh co hai nghiem la x1 = %v va x2 = %v\n", x1, x2)
	}
}
