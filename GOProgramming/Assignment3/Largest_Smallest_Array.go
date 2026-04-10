package main

import "fmt"

func main() {
	arr := []int{10, 25, 5, 50, 15}
	largest := arr[0]
	smallest := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
		if arr[i] < smallest {
			smallest = arr[i]
		}
	}

	fmt.Println("Largest:", largest)
	fmt.Println("Smallest:", smallest)
}
