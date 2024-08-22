package main

import (
	"fmt"

func main() {
	book := service.Book{
		ID:     1,
		Title:  "Harry potter",
		Author: "J.R.R. Tolkien",
		Genre:  "Fantasy",
	}
	fmt.Println("the book is:", book.Title)
}
