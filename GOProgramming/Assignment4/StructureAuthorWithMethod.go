package main

import "fmt"

type Author struct {
	name string
	age  int
}

func (a Author) show() {
	fmt.Println("Name:", a.name)
	fmt.Println("Age:", a.age)
}

func main() {
	a := Author{"Pratiksha", 21}
	a.show()
}
