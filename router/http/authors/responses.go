package authors

import (
	"time"
)

// AuthorResponse struct defines response fields
type AuthorResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ListResponse struct defines authors list response structure
type ListResponse struct {
	Data []AuthorResponse `json:"data"`
}
