package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name  string
	Marks int
}

func main() {
	students := []Student{
		{"Amit", 75},
		{"Neha", 85},
		{"Ravi", 65},
	}

	// Sorting based on marks
	sort.Slice(students, func(i, j int) bool {
		return students[i].Marks < students[j].Marks
	})

	fmt.Println("Sorted Students:")
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i].Name, students[i].Marks)
	}
}
