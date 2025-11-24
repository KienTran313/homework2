package main

import "fmt"

func main() {
	for number := 100; number < 1000; number++ {
		digit1 := number / 100
		digit2 := (number / 10) % 10
		digit3 := number % 10

		total := (digit1 * digit1 * digit1) + (digit2 * digit2 * digit2) + (digit3 * digit3 * digit3)
		if total == number {
			fmt.Println("Cac so armstrong la", number)
		}
	}
}
