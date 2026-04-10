package main

import "fmt"

func operations(a int, b int) (int, int) {
	return a + b, a * b
}

func main() {
	sum, mul := operations(5, 3)

	fmt.Println("Sum:", sum)
	fmt.Println("Multiplication:", mul)
}
