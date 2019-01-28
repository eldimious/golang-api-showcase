package books

import (
	"time"

	authorSchema "github.com/eldimious/golang-api-showcase/data/authors"
)

// struct defines the database model for a Author.
type Book struct {
	ID        int `gorm:"primary_key;type:int;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Publisher string
	Author    authorSchema.Author
	AuthorID  int `gorm:"type:int;"`
}
