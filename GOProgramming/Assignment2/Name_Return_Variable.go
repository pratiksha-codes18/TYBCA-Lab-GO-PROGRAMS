package main

import "fmt"

func calc(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}

func main() {
	s, d := calc(10, 5)
	fmt.Println("Sum:", s)
	fmt.Println("Difference:", d)
}
