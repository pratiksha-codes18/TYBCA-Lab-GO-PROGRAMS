package main

import "fmt"

type Numbers struct {
	a, b int
}

func (n Numbers) Multiply() int {
	return n.a * n.b
}

func main() {
	n := Numbers{3, 4}
	fmt.Println("Multiplication:", n.Multiply())
}
