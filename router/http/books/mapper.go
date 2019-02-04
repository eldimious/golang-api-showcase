package books

import (
	domain "github.com/eldimious/golang-api-showcase/domain/books"
)

func toResponseModel(entity *domain.Book) *BookResponse {
	return &BookResponse{
		Id:        entity.Id,
		Name:      entity.Name,
		AuthorId:  entity.AuthorId,
		Publisher: entity.Publisher,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

}
