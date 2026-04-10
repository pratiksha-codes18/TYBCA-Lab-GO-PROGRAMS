package main

import "fmt"

func main() {
	s1 := "Hello "
	s2 := "World"

	p1 := &s1
	p2 := &s2

	result := *p1 + *p2
	fmt.Println("Concatenated:", result)
}
