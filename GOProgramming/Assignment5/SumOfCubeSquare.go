package main

import "fmt"

func square(n int, ch chan int) {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n = n / 10
	}
	ch <- sum
}

func cube(n int, ch chan int) {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d * d
		n = n / 10
	}
	ch <- sum
}

func main() {
	num := 123
	sqCh := make(chan int)
	cuCh := make(chan int)

	go square(num, sqCh)
	go cube(num, cuCh)

	sq := <-sqCh
	cu := <-cuCh

	fmt.Println("Sum of squares =", sq)
	fmt.Println("Sum of cubes =", cu)
	fmt.Println("Final sum =", sq+cu)
}
