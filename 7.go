package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int
	for {
		fmt.Print("Nhap 1 so nguyen n : ")
		scanner.Scan()
		text := scanner.Text()
		number, err := strconv.Atoi(text)
		n = number
		if err != nil {
			fmt.Println("Loi,chi nhap so nguyen")
		} else {
			break
		}
	}

	var S1 int
	for s := 0; s <= n; s++ {
		S1 += (2*s + 1)
	}
	fmt.Printf("Tong cua S1 = %d\n", S1)

	var S2 float64
	for s := 1; s <= n; s++ {
		S2 += 0.5 / float64(s)
	}
	fmt.Printf("Tong cua S2 = %v\n", S2)

	var S3 int
	for s := 1; s <= n; s++ {
		S3 += (s * (s + 1)) / 2
	}
	fmt.Printf("Tong cua S3 = %d\n", S3)
}
