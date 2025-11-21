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
	var p float64
	for {
		fmt.Print("Nhap so tien muon gui P : ")
		scanner.Scan()
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		p = number
		if err != nil {
			fmt.Println("Loi,hay nhap so.")
		} else {
			break
		}
	}

	var n float64
	for {
		fmt.Print("Nhap so thang gui N :")
		scanner.Scan()
		text := scanner.Text()
		number, err := strconv.ParseFloat(text, 64)
		n = number
		if err != nil {
			fmt.Println("Loi,hay nhap so.")
		} else {
			break
		}
	}
	r := 0.7
	total := p * math.Pow((1+r/100), float64(n))
	// van chua hieu lam day tim hieu ki hon.
	fmt.Printf("Sau %.0f thang, so tien trong tai khoan la: %.0f\n", n, total)
}
