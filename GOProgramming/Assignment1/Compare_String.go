package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Print("Enter first string: ")
	fmt.Scan(&s1)
	fmt.Print("Enter second string: ")
	fmt.Scan(&s2)

	if s1 == s2 {
		fmt.Println("Strings are equal")
	} else {
		fmt.Println("Strings are not equal")
	}
}
