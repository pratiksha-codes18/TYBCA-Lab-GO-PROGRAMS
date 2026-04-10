package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{50, 20, 30, 10, 40}
	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}
