package main

import "fmt"

func main() {
	x := 10
	p := &x
	pp := &p

	fmt.Println("Value of x:", x)
	fmt.Println("Pointer p:", p)
	fmt.Println("Pointer to pointer pp:", pp)
	fmt.Println("Value using pp:", **pp)
}
