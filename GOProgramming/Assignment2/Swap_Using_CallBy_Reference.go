package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x := 10
	y := 20

	swap(&x, &y)

	fmt.Println("After swap:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
}
