package books

import (
	"time"
)

// BookResponse struct defines response fields
type BookResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Publisher string    `json:"surname"`
	AuthorId  int       `json:"authorId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ListResponse struct defines books list response structure
type ListResponse struct {
	Data []BookResponse `json:"data"`
}
