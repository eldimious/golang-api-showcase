package books

// BookRepository provides an abstraction on top of the book data source
type BookRepository interface {
	CreateBook(*Book) (*Book, error)
	ReadBook(int) (*Book, error)
	ListBooks(authorId int) ([]Book, error)
}
