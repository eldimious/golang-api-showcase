package authors

import (
	"time"
)

// struct defines the database model for a Author.
type Author struct {
	ID        int `gorm:"primary_key;type:int;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Surname   string
}
