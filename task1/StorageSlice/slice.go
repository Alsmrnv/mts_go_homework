package StorageSlice

import "LibraryProject/book"

type BookSlice struct {
	books []book.Book
}

func (s *BookSlice) GetBook(id uint, title string) (book.Book, bool) {
	for _, curBook := range s.books {
		if curBook.Title == title {
			return curBook, true
		}
	}
	return book.Book{}, false
}

func (s *BookSlice) AddBook(newBook *book.Book) {
	s.books = append(s.books, *newBook)
}

func (s *BookSlice) Rehash(newGenID func(string) uint) {
	for _, curBook := range s.books {
		curBook.Id = newGenID(curBook.Title) // TODO if its value, not ptr then problems
	}
}

func (s *BookSlice) GetAllBooks() []book.Book {
	return s.books
}

func (s *BookSlice) AddAllBooks(books []book.Book) {
	s.books = books
}
