package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var number int
	var sum int
	var even int
	for {
		fmt.Print("Nhap cac so nguyen [1-10] va (nhap 0 de dung lai): ")
		scanner.Scan()
		text := scanner.Text()
		a, err := strconv.Atoi(text)
		number = a
		if err != nil {
			fmt.Println("Loi,hay nhap so [1-10].")
		} else if number == 0 {
			break
		} else if number < 1 || number > 10 {
			fmt.Println("Loi,chi nhap so trong khoang [1-10].")
		} else if number%2 == 0 {
			even = even + 1
		} else {
			sum += number
		}
	}
	fmt.Printf("Tong cac so chan la : %d\n", even)
	fmt.Printf("Tong cua cac so le la : %d\n", sum)
}
