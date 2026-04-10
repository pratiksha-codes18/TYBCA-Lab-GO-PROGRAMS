package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1, s2 string

	fmt.Print("Enter main string: ")
	fmt.Scan(&s1)
	fmt.Print("Enter substring: ")
	fmt.Scan(&s2)

	if strings.Contains(s1, s2) {
		fmt.Println("Substring found")
	} else {
		fmt.Println("Substring not found")
	}
}
