package authors

// AuthorRepository provides an abstraction on top of the author data source
type AuthorRepository interface {
	CreateAuthor(*Author) (*Author, error)
	ReadAuthor(int) (*Author, error)
	ListAuthors() ([]Author, error)
}
