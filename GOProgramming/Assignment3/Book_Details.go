package main

import "fmt"

type Book struct {
	BookID int
	Title  string
	Author string
	Price  float64
}

func main() {
	var n int
	fmt.Print("Enter number of books: ")
	fmt.Scan(&n)

	books := make([]Book, n)

	
	for i := 0; i < n; i++ {
		fmt.Print("BookID: ")
		fmt.Scan(&books[i].BookID)
		fmt.Print("Title: ")
		fmt.Scan(&books[i].Title)
		fmt.Print("Author: ")
		fmt.Scan(&books[i].Author)
		fmt.Print("Price: ")
		fmt.Scan(&books[i].Price)
	}

	
	fmt.Println("\nBook Details:")
	for i := 0; i < n; i++ {
		fmt.Println("ID:", books[i].BookID,
			"Title:", books[i].Title,
			"Author:", books[i].Author,
			"Price:", books[i].Price)
	}
}
