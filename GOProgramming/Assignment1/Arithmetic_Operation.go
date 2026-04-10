package main

import "fmt"

func main() {
	var a, b, choice int

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)

	fmt.Println("1.Add 2.Subtract 3.Multiply 4.Divide")
	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Addition:", a+b)
	case 2:
		fmt.Println("Subtraction:", a-b)
	case 3:
		fmt.Println("Multiplication:", a*b)
	case 4:
		fmt.Println("Division:", a/b)
	default:
		fmt.Println("Invalid choice")
	}
}
