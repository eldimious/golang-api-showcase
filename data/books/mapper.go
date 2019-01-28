package books

import (
	domain "github.com/eldimious/golang-api-showcase/domain/books"
)

func toDBModel(entity *domain.Book) *Book {
	return &Book{
		ID:        entity.ID,
		Name:      entity.Name,
		Publisher: entity.Publisher,
		AuthorID:  entity.AuthorID,
	}
}

func toDomainModel(entity *Book) *domain.Book {
	return &domain.Book{
		ID:        entity.ID,
		Name:      entity.Name,
		Publisher: entity.Publisher,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}
