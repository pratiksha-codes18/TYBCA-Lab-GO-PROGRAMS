package main

import "fmt"

type Printer interface {
	print()
}

type Scanner interface {
	scan()
}

// Embedded Interface
type Machine interface {
	Printer
	Scanner
}

type Device struct{}

func (d Device) print() {
	fmt.Println("Printing...")
}

func (d Device) scan() {
	fmt.Println("Scanning...")
}

func main() {
	var m Machine
	m = Device{}

	m.print()
	m.scan()
}
