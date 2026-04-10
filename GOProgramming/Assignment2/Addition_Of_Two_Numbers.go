package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	var x, y int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x, &y)

	result := add(x, y)
	fmt.Println("Addition:", result)
}
