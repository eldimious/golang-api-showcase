package authors

import (
	"time"
)

// struct defines the database model for a Author.
type Author struct {
	Id        int `gorm:"primary_key";"AUTO_INCREMENT";`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
	Surname   string
}
