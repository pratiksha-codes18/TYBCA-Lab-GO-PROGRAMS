package main

import "fmt"

func main() {
	numbers := []int{
		10,
		20,
		30,
		40,
		50,
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element:", numbers[i])
	}
}
