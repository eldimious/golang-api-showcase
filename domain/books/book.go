package books

import (
	"time"
)

//struct contains information about a book.
type Book struct {
	ID        int
	Name      string
	Publisher string
	AuthorID  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
