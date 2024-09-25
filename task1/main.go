package main

import (
	library2 "LibraryProject/library"
	"fmt"
)

func newGenID(title string) uint {
	var sum uint = 0
	for _, char := range title {
		sum += uint(char)*2 + 1
	}
	return sum
}

func main() {
	library := library2.NewLibrary()
	books := make([]library2.Book, 0)
	books = append(books,
		library2.NewBook("Harry Potter", "Joanne Rowling", 1997),
		library2.NewBook("The Little Prince", "Antoine de Saint-Exupery", 1943),
		library2.NewBook("The Alchemist", "Paulo Coelho", 1988),
		library2.NewBook("And Then There Were None", "Agatha Christie", 1939),
		library2.NewBook("The Hobbit", "John Ronald Reuel Tolkien", 1937),
		library2.NewBook("Alice's Adventures in Wonderland", "Lewis Carroll", 1865))
	for _, book := range books {
		library.AddBook(book)
	}

	fmt.Println("---------------------------\nBefore changes:\n---------------------------\n")

	book1, ok1 := library.GetBook("Harry Potter")
	if ok1 == false {
		fmt.Printf("Harry Potter not found")
	} else {
		fmt.Printf("Book: '%s'\nAuthor: %s\nYear: %d\n\n", book1.Title, book1.Author, book1.PublicationYear)
	}

	book2, ok2 := library.GetBook("War and peace")
	if ok2 == false {
		fmt.Printf("'War and peace' not found\n\n")
	} else {
		fmt.Printf("Book: '%s'\nAuthor: %s\nYear: %d\n\n", book2.Title, book2.Author, book2.PublicationYear)
	}

	library.ChangeGeneratorID(newGenID)

	fmt.Println("---------------------------\nAfter changing id generator:\n---------------------------\n")

	book3, ok3 := library.GetBook("Alice's Adventures in Wonderland")
	if ok3 == false {
		fmt.Printf("'Alice's Adventures in Wonderland' not found")
	} else {
		fmt.Printf("Book: '%s'\nAuthor: %s\nYear: %d\n\n", book3.Title, book3.Author, book3.PublicationYear)
	}

	var newStorage library2.BookSlice
	library.ChangeStorage(&newStorage)

	fmt.Println("---------------------------\nAfter changing storage:\n---------------------------\n")

	book4, ok4 := library.GetBook("The Hobbit")
	if ok4 == false {
		fmt.Printf("'The Hobbit' not found")
	} else {
		fmt.Printf("Book: '%s'\nAuthor: %s\nYear: %d\n\n", book4.Title, book4.Author, book4.PublicationYear)
	}

}
