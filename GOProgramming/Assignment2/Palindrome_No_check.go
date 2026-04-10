package main

import "fmt"

func isPalindrome(n int) bool {
	original := n
	rev := 0

	for n > 0 {
		rem := n % 10
		rev = rev*10 + rem
		n = n / 10
	}

	return original == rev
}

func main() {
	var num int
	fmt.Print("Enter number: ")
	fmt.Scan(&num)

	if isPalindrome(num) {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}
