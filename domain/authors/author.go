package authors

import (
	"time"
)

// Author struct contains information an author.
type Author struct {
	Id        int
	Name      string
	Surname   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
