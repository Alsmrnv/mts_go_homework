package main

import (
	"LibraryProject/book"
	"LibraryProject/library"
	"LibraryProject/storageslice"
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
	lib := library.NewLibrary()
	books := make([]book.Book, 0)
	books = append(books,
		book.NewBook("Harry Potter", "Joanne Rowling", 1997),
		book.NewBook("The Little Prince", "Antoine de Saint-Exupery", 1943),
		book.NewBook("The Alchemist", "Paulo Coelho", 1988),
		book.NewBook("And Then There Were None", "Agatha Christie", 1939),
		book.NewBook("The Hobbit", "John Ronald Reuel Tolkien", 1937),
		book.NewBook("Alice's Adventures in Wonderland", "Lewis Carroll", 1865))
	for _, curBook := range books {
		lib.AddBook(curBook)
	}

	fmt.Println("---------------------------\nBefore changes:\n---------------------------\n")

	book1, ok1 := lib.GetBook("Harry Potter")
	if ok1 == false {
		fmt.Printf("Harry Potter not found")
	} else {
		fmt.Printf("Book: '%s'\nAuthor: %s\nYear: %d\n\n", book1.Title, book1.Author, book1.PublicationYear)
	}

	book2, ok2 := lib.GetBook("War and peace")
	if ok2 == false {
		fmt.Printf("'War and peace' not found\n\n")
	} else {
		fmt.Printf("Book: '%s'\nAuthor: %s\nYear: %d\n\n", book2.Title, book2.Author, book2.PublicationYear)
	}

	lib.ChangeGeneratorID(newGenID)

	fmt.Println("---------------------------\nAfter changing id generator:\n---------------------------\n")

	book3, ok3 := lib.GetBook("Alice's Adventures in Wonderland")
	if ok3 == false {
		fmt.Printf("'Alice's Adventures in Wonderland' not found")
	} else {
		fmt.Printf("Book: '%s'\nAuthor: %s\nYear: %d\n\n", book3.Title, book3.Author, book3.PublicationYear)
	}

	var newStorage storageslice.BookSlice
	lib.ChangeStorage(&newStorage)

	fmt.Println("---------------------------\nAfter changing storage:\n---------------------------\n")

	book4, ok4 := lib.GetBook("The Hobbit")
	if ok4 == false {
		fmt.Printf("'The Hobbit' not found")
	} else {
		fmt.Printf("Book: '%s'\nAuthor: %s\nYear: %d\n\n", book4.Title, book4.Author, book4.PublicationYear)
	}

}
