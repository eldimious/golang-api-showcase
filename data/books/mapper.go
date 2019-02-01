package books

import (
	domain "github.com/eldimious/golang-api-showcase/domain/books"
)

func toDBModel(entity *domain.Book) *Book {
	return &Book{
		Id:        entity.Id,
		Name:      entity.Name,
		Publisher: entity.Publisher,
		AuthorId:  entity.AuthorId,
	}
}

func toDomainModel(entity *Book) *domain.Book {
	return &domain.Book{
		Id:        entity.Id,
		Name:      entity.Name,
		Publisher: entity.Publisher,
		AuthorId:  entity.AuthorId,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}
