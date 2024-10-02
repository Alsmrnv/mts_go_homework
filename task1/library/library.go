package library

import (
	"LibraryProject/book"
	"LibraryProject/storage"
	"LibraryProject/storagemap"
)

type Library struct {
	bookStorage storage.Storage
	genID       func(string) uint
}

func NewLibrary() *Library {
	return &Library{storagemap.NewBookMap(), book.HashBook}
}

func (lib *Library) GetBook(title string) (book.Book, bool) {
	return lib.bookStorage.GetBook(lib.genID(title), title)
}

func (lib *Library) AddBook(newBook book.Book) { // TODO better ptr, but implicit cast to ptr works ????
	_, ok := lib.GetBook(newBook.Title)
	if ok {
		return
	}

	id := lib.genID(newBook.Title)
	newBook.Id = id

	lib.bookStorage.AddBook(&newBook)
}

func (lib *Library) ChangeGeneratorID(newGenID func(string) uint) {
	lib.genID = newGenID
	lib.bookStorage.Rehash(lib.genID)
}

func (lib *Library) ChangeStorage(s storage.Storage) {
	saveBooks := lib.bookStorage.GetAllBooks()
	s.AddAllBooks(saveBooks)
	lib.bookStorage = s
}
