package main

import (
	"fmt"
	"time"
)

func printNumbers(id int) {
	for i := 0; i <= 10; i++ {
		fmt.Println("Goroutine", id, ":", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	for i := 1; i <= 5; i++ {
		go printNumbers(i)
	}

	time.Sleep(time.Second * 2)
}
