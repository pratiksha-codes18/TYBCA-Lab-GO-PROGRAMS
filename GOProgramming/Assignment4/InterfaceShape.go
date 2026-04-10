package main

import "fmt"

// Interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// Rectangle
type Rectangle struct {
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	c := Circle{5}
	r := Rectangle{4, 6}

	fmt.Println("Circle Area:", c.Area())
	fmt.Println("Circle Perimeter:", c.Perimeter())

	fmt.Println("Rectangle Area:", r.Area())
	fmt.Println("Rectangle Perimeter:", r.Perimeter())
}
