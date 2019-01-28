package books

// BookRepository provides an abstraction on top of the book data source
type BookRepository interface {
	CreateBook(*Book) (*Book, error)
	ReadBook(int) (*Book, error)
	ListBooks(authorID int) ([]Book, error)
}
