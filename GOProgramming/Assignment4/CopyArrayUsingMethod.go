package main

import "fmt"

func copyArr(src [5]int) [5]int {
	var dest [5]int

	for i := 0; i < 5; i++ {
		dest[i] = src[i]
	}

	return dest
}

func main() {
	src := [5]int{1, 2, 3, 4, 5}

	dest := copyArr(src)

	fmt.Println("Source:", src)
	fmt.Println("Copied:", dest)
}
