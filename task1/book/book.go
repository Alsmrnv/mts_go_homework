package book

type Book struct {
	Title           string
	Author          string
	PublicationYear int
	Id              uint
}

func NewBook(title string, author string, year int) Book {
	return Book{title, author, year, 0}
}

func HashBook(title string) uint {
	var sum uint = 0
	for _, char := range title {
		sum += uint(char)
	}
	return sum
}
