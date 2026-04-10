package main

import "fmt"

type Student struct {
	name string
	roll int
}

func (s *Student) show() {
	fmt.Println("Name:", s.name)
	fmt.Println("Roll No:", s.roll)
}

func main() {
	s := Student{"Rahul", 101}
	s.show()
}
