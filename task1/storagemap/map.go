package storagemap

import "LibraryProject/book"

type BookMap struct {
	books map[uint][]book.Book
}

func (s *BookMap) GetBook(id uint, title string) (book.Book, bool) {
	books, ok := s.books[id]
	if ok == false {
		return book.Book{}, false
	}

	for _, curBook := range books {
		if curBook.Title == title {
			return curBook, true
		}
	}

	return book.Book{}, false
}

func (s *BookMap) AddBook(newBook *book.Book) {
	s.books[newBook.Id] = append(s.books[newBook.Id], *newBook)
}

func (s *BookMap) Rehash(newGenID func(string) uint) {
	for _, bookSlice := range s.books {
		for i, curBook := range bookSlice {
			curBook.Id = newGenID(curBook.Title)
			bookSlice = append(bookSlice[:i], bookSlice[i+1:]...)
			s.books[curBook.Id] = append(s.books[curBook.Id], curBook)
		}
	}
}

func (s *BookMap) GetAllBooks() []book.Book {
	allBooks := make([]book.Book, 0)

	for _, bookSlice := range s.books {
		for _, curBook := range bookSlice {
			allBooks = append(allBooks, curBook)
		}
	}

	return allBooks
}

func (s *BookMap) AddAllBooks(books []book.Book) {
	for _, curBook := range books {
		s.books[curBook.Id] = append(s.books[curBook.Id], curBook)
	}
}

func NewBookMap() *BookMap {
	newMap := make(map[uint][]book.Book)
	return &BookMap{newMap}
}
