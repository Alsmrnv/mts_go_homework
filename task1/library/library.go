package library

type Library struct {
	storage Storage
	genID   func(string) uint
}

type Book struct {
	Title           string
	Author          string
	PublicationYear int
	id              uint
}

type Storage interface {
	GetBook(id uint, title string) (Book, bool)
	AddBook(book *Book)
	Rehash(newGenID func(string) uint)
	GetAllBooks() []Book
	AddAllBooks(books []Book)
}

func NewLibrary() *Library {
	return &Library{&BookMap{make(map[uint][]Book)}, HashBook}
}

func NewBook(title string, author string, year int) Book {
	return Book{title, author, year, 0}
}

func (lib *Library) GetBook(title string) (Book, bool) {
	return lib.storage.GetBook(lib.genID(title), title)
}

func (lib *Library) AddBook(book Book) { // TODO better ptr, but implicit cast to ptr works ????
	_, ok := lib.GetBook(book.Title)
	if ok {
		return
	}

	id := lib.genID(book.Title)
	book.id = id

	lib.storage.AddBook(&book)
}

func (lib *Library) ChangeGeneratorID(newGenID func(string) uint) {
	lib.genID = newGenID
	lib.storage.Rehash(lib.genID)
}

func (lib *Library) ChangeStorage(s Storage) {
	saveBooks := lib.storage.GetAllBooks()
	s.AddAllBooks(saveBooks)
	lib.storage = s
}

func HashBook(title string) uint {
	var sum uint = 0
	for _, char := range title {
		sum += uint(char)
	}
	return sum
}

type BookSlice struct {
	books []Book
}

type BookMap struct {
	books map[uint][]Book
}

func (s *BookSlice) GetBook(id uint, title string) (Book, bool) {
	for _, book := range s.books {
		if book.Title == title {
			return book, true
		}
	}
	return Book{}, false
}

func (s *BookSlice) AddBook(book *Book) {
	s.books = append(s.books, *book)
}

func (s *BookSlice) Rehash(newGenID func(string) uint) {
	for _, book := range s.books {
		book.id = newGenID(book.Title) // TODO if its value, not ptr then problems
	}
}

func (s *BookSlice) GetAllBooks() []Book {
	return s.books
}

func (s *BookSlice) AddAllBooks(books []Book) {
	s.books = books
}

func (s *BookMap) GetBook(id uint, title string) (Book, bool) {
	books, ok := s.books[id]
	if ok == false {
		return Book{}, false
	}

	for _, book := range books {
		if book.Title == title {
			return book, true
		}
	}

	return Book{}, false
}

func (s *BookMap) AddBook(book *Book) {
	s.books[book.id] = append(s.books[book.id], *book)
}

func (s *BookMap) Rehash(newGenID func(string) uint) {
	for _, bookSlice := range s.books {
		for i, book := range bookSlice {
			book.id = newGenID(book.Title)
			bookSlice = append(bookSlice[:i], bookSlice[i+1:]...)
			s.books[book.id] = append(s.books[book.id], book)
		}
	}
}

func (s *BookMap) GetAllBooks() []Book {
	allBooks := make([]Book, 0)

	for _, bookSlice := range s.books {
		for _, book := range bookSlice {
			allBooks = append(allBooks, book)
		}
	}

	return allBooks
}

func (s *BookMap) AddAllBooks(books []Book) {
	for _, book := range books {
		s.books[book.id] = append(s.books[book.id], book)
	}
}
