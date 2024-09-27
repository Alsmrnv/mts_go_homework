package storage

import "LibraryProject/book"

type Storage interface {
	GetBook(id uint, title string) (book.Book, bool)
	AddBook(book *book.Book)
	Rehash(newGenID func(string) uint)
	GetAllBooks() []book.Book
	AddAllBooks(books []book.Book)
}
