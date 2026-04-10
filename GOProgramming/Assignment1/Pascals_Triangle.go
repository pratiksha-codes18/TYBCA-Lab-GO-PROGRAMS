package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter rows: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		num := 1
		for j := 0; j <= i; j++ {
			fmt.Print(num, " ")
			num = num * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}
