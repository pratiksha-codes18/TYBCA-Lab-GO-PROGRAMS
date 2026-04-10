package main

import "fmt"

func change(a int) {
	a = 100
}

func main() {
	x := 10
	change(x)

	fmt.Println("Value of x:", x) 
}
