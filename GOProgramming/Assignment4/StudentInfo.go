package main

import "fmt"

type Student struct {
	roll int
	name string
	per  float64
}

func main() {
	students := []Student{
		{1, "A", 75},
		{2, "B", 85},
		{3, "C", 65},
	}

	// Sorting (simple bubble sort)
	for i := 0; i < len(students); i++ {
		for j := i + 1; j < len(students); j++ {
			if students[i].per < students[j].per {
				students[i], students[j] = students[j], students[i]
			}
		}
	}

	fmt.Println("Students in Descending Order:")
	for _, s := range students {
		fmt.Println(s.roll, s.name, s.per)
	}
}
