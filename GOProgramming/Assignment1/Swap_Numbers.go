package main

import "fmt"

func main() {
	a := 5
	b := 10
	fmt.Println("Before swap:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	a = a + b
	b = a - b
	a = a - b

	fmt.Println("After swap:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
