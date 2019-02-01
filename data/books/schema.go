package books

import (
	"time"

	authorSchema "github.com/eldimious/golang-api-showcase/data/authors"
)

// struct defines the database model for a Author.
type Book struct {
	Id        int `gorm:"primary_key";"AUTO_INCREMENT";`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Publisher string
	Author    authorSchema.Author
	AuthorId  int `gorm:"type:int;"`
}
