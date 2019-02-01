package books

import (
	"time"
)

//struct contains information about a book.
type Book struct {
	Id        int
	Name      string
	Publisher string
	AuthorId  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
